/*
@Time : 2019-01-17 17:01 
@Author : vannnnish
@File : common
*/

package apppay

import (
	"bytes"
	"encoding/xml"
	"github.com/vannnnish/yeego/yeerand"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// 下单接口
	getPrepayUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	// 查询支付结果
	payResultUrl = "https://api.mch.weixin.qq.com/pay/orderquery"
)

// app下单
/*
	body:       这条交易的的信息   例如: 王者荣耀增值服务
	outTradeNo: 你自己服务器上的订单号(自己生成的，且不能重复)
	totalFee:   总金额，单位是分
	notifyUrl:  支付完成后的回调，你自己服务器提供的接口，
	tradeType:  APP
	key:        微信商户号上面设置的apiKey
*/
func GenerateWechatOrder(appId, mchId, nonceStr, body, outTradeNo string, totalFee int, spbillCreateIp, notifyUrl, tradeType, key string) (WechatResponse, error) {
	wechatResponse := WechatResponse{}
	// 微信统一下单
	wechatOrder := &AppSignStruct{
		// 必填
		AppId:          appId,
		MchId:          mchId,
		NonceStr:       nonceStr,
		Body:           body,
		OutTradeNo:     outTradeNo,
		TotalFee:       totalFee,
		SpbillCreateIp: spbillCreateIp,
		NotifyUrl:      notifyUrl,
		TradeType:      tradeType,
	}
	out := dataToMap(*wechatOrder)
	wechatOrder.Sign = Sign(out, key)
	buf, err := xml.Marshal(wechatOrder)
	if err != nil {
		return wechatResponse, nil
	}
	// 将订单对象转化为xml
	tmp := strings.Replace(string(buf), "<AppSignStruct>", "<xml>", -1)
	tmp = strings.Replace(tmp, "</AppSignStruct>", "</xml>", -1)
	// 转化后的字符串
	bufData := bytes.NewBuffer([]byte(tmp))
	// 发送请求
	r, err := http.Post(getPrepayUrl, "text/xml", bufData)
	if err != nil {
		return wechatResponse, err
	}
	response, err := ioutil.ReadAll(r.Body)

	err = xml.Unmarshal(response, &wechatResponse)
	if err != nil {
		return wechatResponse, err
	}
	return wechatResponse, nil
}

// 支付查询
// outTradeNo: 商户自己生成的订单号 (也可以传微信的订单号，但是觉得自己的就好了)
// 自己判断QueryWechatResponse.Result_code 如果是SUCCESS 表示支付成功，否则失败了
func QueryPayResult(appId, mchId, outTradeNo string, key string) (QueryWechatResponse, error) {
	var res QueryWechatResponse
	nonce := yeerand.RandString(32)
	query := QueryWechatRequest{
		Appid:        appId,
		Mch_id:       mchId,
		Out_trade_no: outTradeNo,
		Nonce_str:    nonce,
	}
	dataMap := dataToMap(query)
	sign := Sign(dataMap, key)
	query.Sign = sign
	buf, err := xml.Marshal(&query)
	if err != nil {
		return res, err
	}
	// 将订单对象转化为xml
	tmp := strings.Replace(string(buf), "<QueryWechatRequest>", "<xml>", -1)
	tmp = strings.Replace(tmp, "</QueryWechatRequest>", "</xml>", -1)
	// 转化后的字符串
	body := bytes.NewBuffer([]byte(tmp))
	r, err := http.Post(payResultUrl, "text/xml", body)
	if err != nil {
		return res, err
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return res, err
	}
	err = xml.Unmarshal(response, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
