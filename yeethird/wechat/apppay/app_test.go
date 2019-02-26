/*
@Time : 2019-02-18 15:01 
@Author : vannnnish
@File : common_test
*/

package apppay

import (
	"fmt"
	"github.com/vannnnish/yeego"
	"testing"
)

var (
	appId = ""
	mchId = ""
	key   = ""
)

func Init() {
	yeego.MustInitConfig("/Users/vannnnish/go_project/project_sh/yeego/conf", "conf")
	appId = yeego.Config.GetString("wechat.AppId")
	mchId = yeego.Config.GetString("wechat.MchId")
	key = yeego.Config.GetString("wechat.Key")
}

// 测试app支付下单
func TestGenerateWechatOrder(t *testing.T) {
	Init()
	nonceStr := "sdfsdf"
	body := "你好"
	outTradeNo := "No1234431122324"
	totalFee := 1
	spbillCreateIp := "10.1.1.111"
	notifyUrl := "http://www.baidu.com"
	tradeType := "APP"
	prePay, err := GenerateWechatOrder(appId, mchId, nonceStr, body, outTradeNo, totalFee, spbillCreateIp, notifyUrl, tradeType, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prePay)
}

func TestQueryWechat(t *testing.T) {
	Init()
	outTradeNum := "aidian20190223003430gizc"
	res, err := QueryPayResult(appId, mchId, outTradeNum, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("结果:", res)
}
