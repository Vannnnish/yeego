/**
 * Created by angelina on 2017/7/31.
 * https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_1
 */

package yeeWechat

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"github.com/labstack/echo"
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeRand"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	// 支付成功回调地址
	PaySuccessCallbackUrl = "/Wx/PayCallback"
)

// 微信支付实例
type Wx struct {
	AppId                string                                           // 公众账号ID
	AppSecret            string                                           // 应用密钥
	MchId                string                                           // 绑定商户ID
	Key                  string                                           // 微信商户平台：账户设置-安全设置-API安全-API密钥
	BaseAuthCallBack     func(req *yeego.Request) error                   // Base授权回调
	UserInfoAuthCallBack func(req *yeego.Request) error                   // 高级授权回调
	PaySuccessCallBack   func(req *yeego.Request, payInfo PaySuccessInfo) // 支付回调
}

// 微信统一下单数据结构
type WxOrder struct {
	AppId          string `xml:"appid"`            //
	Body           string `xml:"body"`             // 商品或支付单简要描述
	MchId          string `xml:"mch_id"`           // 商户号
	NonceStr       string `xml:"nonce_str"`        // 随机字符串
	NotifyUrl      string `xml:"notify_url"`       // 接收微信支付异步通知回调地址
	Openid         string `xml:"openid"`           //
	OutTradeNo     string `xml:"out_trade_no"`     // 商户订单号
	Sign           string `xml:"sign"`             // 签名
	SpBillCreateIp string `xml:"spbill_create_ip"` // 终端IP
	TotalFee       string `xml:"total_fee"`        // 总金额(分)
	TradeType      string `xml:"trade_type"`       // 交易类型 取 JSAPI
	Attach         string `xml:"attach"`           // 附加数据
}

// 统一下单返回数据
type WxResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	PrePayId   string `xml:"prepay_id"` // 订单Id
}

// 前台JS下单数据结构
type Order struct {
	OpenId     string
	Body       string // 商品描述
	OutTradeNo string // 商户系统的订单号，与请求一致
	TotalFee   string // 订单总金额，单位为分
	Attach     string // 商家数据包，原样返回
}

// JS下单返回数据结构
type PreOrder struct {
	AppId     string
	TimeStamp string
	NonceStr  string
	SignType  string
	Package   string
	PaySign   string
}

// 生成JS调用订单
func (w *Wx) GenerateWxOrder(c echo.Context, order *Order) PreOrder {
	// 微信统一下单
	wxOrder := &WxOrder{
		AppId:          w.AppId,
		MchId:          w.MchId,
		NonceStr:       yeeRand.RandString(16),
		NotifyUrl:      yeego.Config.GetString("SelfSchemeAndDomain") + PaySuccessCallbackUrl,
		Openid:         order.OpenId,
		Body:           order.Body,
		OutTradeNo:     order.OutTradeNo,
		TotalFee:       order.TotalFee,
		SpBillCreateIp: c.RealIP(),
		TradeType:      "JSAPI",
		Attach:         order.Attach,
	}
	out := dataToMap(*wxOrder)
	wxOrder.Sign = w.sign(out)
	buf, _ := xml.Marshal(wxOrder)
	tmp := strings.Replace(string(buf), "<WechatOrder>", "", -1)
	tmp = strings.Replace(tmp, "</WechatOrder>", "", -1)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	wxResponse := WxResponse{}
	xml.Unmarshal(response, &wxResponse)
	if wxResponse.ReturnCode != "SUCCESS" {
		panic(wxResponse.ReturnMsg)
	}
	preOrder := PreOrder{
		AppId:     w.AppId,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  yeeRand.RandString(16),
		SignType:  "MD5",
		Package:   "prepay_id=" + wxResponse.PrePayId,
	}
	preOrder.PaySign = w.sign(map[string]string{
		"appId":     preOrder.AppId,
		"timeStamp": preOrder.TimeStamp,
		"nonceStr":  preOrder.NonceStr,
		"signType":  preOrder.SignType,
		"package":   preOrder.Package,
	})
	return preOrder
}

// 支付成功返回信息
type PaySuccessInfo struct {
	Appid          string `xml:"appid"`
	Bank_type      string `xml:"bank_type"`      // 银行类型
	Cash_fee       string `xml:"cash_fee"`       // 现金支付金额订单现金支付金额
	Fee_type       string `xml:"fee_type"`       // 货币类型
	Is_subscribe   string `xml:"is_subscribe"`   // 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	Mch_id         string `xml:"mch_id"`         //
	Nonce_str      string `xml:"nonce_str"`      //
	Openid         string `xml:"openid"`         //
	Out_trade_no   string `xml:"out_trade_no"`   //
	Result_code    string `xml:"result_code"`    //
	Return_code    string `xml:"return_code"`    //
	Sign           string `xml:"sign"`           // 签名
	Time_end       string `xml:"time_end"`       // 支付完成时间，格式为yyyyMMddHHmmss
	Total_fee      string `xml:"total_fee"`      // 订单总金额
	Trade_type     string `xml:"trade_type"`     //
	Transaction_id string `xml:"transaction_id"` // 微信支付订单号
	Attach         string `xml:"attach"`         //
}

// 给微信服务器回复的消息接口
type WxReturn struct {
	ReturnCode string `xml:"return_code"`
}

// 支付成功回调
func (w *Wx) PaySuccess() echo.HandlerFunc {
	f := func(c echo.Context) error {
		req, _ := yeego.NewReqAndRes(c)
		out, _ := ioutil.ReadAll(c.Request().Body)
		payInfo := PaySuccessInfo{}
		xml.Unmarshal(out, &payInfo)
		data := dataToMap(payInfo)
		// 验证签名
		if w.sign(data) == payInfo.Sign {
			w.PaySuccessCallBack(req, payInfo)
			// 回复微信服务器SUCCESS
			data := WxReturn{"SUCCESS"}
			x, err := xml.MarshalIndent(data, "", "  ")
			if err != nil {
				panic(err)
			}
			c.Request().Header.Set("Content-Type", "application/xml")
			return c.String(200, string(x))
		} else {
			yeego.DefaultLogError("PaySuccess Sign Error")
			panic("签名错误")
		}
	}
	return f
}

// todo 微信订单查询接口

func dataToMap(obj interface{}) map[string]string {
	v := reflect.ValueOf(obj)
	values := map[string]string{}
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		key := v.Type().Field(i).Name
		tmpKey := ""
		if key != "Sign" && val != "" {
			tmpKey = key
			a := []rune(tmpKey)
			a[0] = unicode.ToLower(a[0])
			tmpKey = string(a)
			values[tmpKey] = val.(string)
		}
	}
	return values
}

func (w *Wx) sign(parameters map[string]string) string {
	ks := make([]string, 0, len(parameters))
	for k := range parameters {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	h := md5.New()
	signature := make([]byte, h.Size()*2)
	for _, k := range ks {
		v := parameters[k]
		if v == "" {
			continue
		}
		h.Write([]byte(k))
		h.Write([]byte{'='})
		h.Write([]byte(v))
		h.Write([]byte{'&'})
	}
	h.Write([]byte("key="))
	h.Write([]byte(w.Key))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}
