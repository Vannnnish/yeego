/*
@Time : 2019-02-18 15:01 
@Author : vannnnish
@File : common_test
*/

package apppay

import (
	"fmt"
	"testing"
)

// 测试app支付下单
func TestGenerateWechatOrder(t *testing.T) {
	appId := "wxf876b6d694fc4ff3"
	mchId := "1524118201"
	nonceStr := "sdfsdf"
	body := "你好"
	outTradeNo := "No1234431122324"
	totalFee := "10"
	key := "LsKE5FFen5dhFEX6wObOFQkdChgKYUGK"
	spbillCreateIp := "10.1.1.111"
	notifyUrl := "http://www.baidu.com"
	tradeType := "APP"
	prePay := GenerateWechatOrder(appId, mchId, nonceStr, body, outTradeNo, totalFee, spbillCreateIp, notifyUrl, tradeType, key)
	fmt.Println(prePay)
}

// 测试app支付下单
func TestGenerateWechatOrderForJs(t *testing.T) {
	//appId := "wxf876b6d694fc4ff3"
	appId := "wxa8d4d9bf3bf8bcda"
	//mchId := "1524118201"
	mchId := "1343373201"
	nonceStr := "sdfsdf"
	body := "你好"
	outTradeNo := "No12344311234"
	totalFee := "100"
	//key := "LsKE5FFen5dhFEX6wObOFQkdChgKYUGK"
	key := "F6bVg9dyrNF6bVg9dyrNF6bVg9dyrNyy1"
	spbillCreateIp := "10.1.1.111"
	notifyUrl := "http://www.baidu.com"
	tradeType := "APP"
	prePay := GenerateWechatOrder(appId, mchId, nonceStr, body, outTradeNo, totalFee, spbillCreateIp, notifyUrl, tradeType, key)
	fmt.Println(prePay)
}
