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
	"strconv"
	"strings"
	"time"
)

// 生成JS调用订单
func GenerateWechatOrder(appId, mchId, nonceStr, body, outTradeNo, totalFee, spbillCreateIp, notifyUrl, tradeType, key string) PreOrder {
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
	// 将订单对象转化为xml
	tmp := strings.Replace(string(buf), "<AppSignStruct>", "<xml>", -1)
	tmp = strings.Replace(tmp, "</AppSignStruct>", "</xml>", -1)
	// 转化后的字符串
	bufData := bytes.NewBuffer([]byte(tmp))
	// 发送请求
	r, err := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "text/xml", bufData)
	if err != nil {
		return PreOrder{}
	}
	response, err := ioutil.ReadAll(r.Body)
	wechatResponse := WechatResponse{}
	xml.Unmarshal(response, &wechatResponse)
	if wechatResponse.ReturnCode != "SUCCESS" {
		panic(wechatResponse.ReturnMsg)
	}
	preOrder := PreOrder{
		AppId:     appId,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  yeerand.RandString(16),
		SignType:  "MD5",
		Package:   "prepay_id=",
	}
	return preOrder
}
