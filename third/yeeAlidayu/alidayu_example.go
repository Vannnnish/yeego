/**
 * Created by angelina on 2017/7/31.
 */

package yeeAlidayu

import (
	"fmt"
	"github.com/yeeyuntech/yeego/third/yeeAlidayu/request"
)

/*
{
	"alibaba_aliqin_fc_sms_num_send_response":
	{
		"result":
		{
		"err_code": "0",
		"model": "106492782347^1108812166203",
		"success": true
		},
	"request_id": "qm2rnyy1ph90"
	}
}
{
	"error_response":
	{
		"code": 15,
		"msg": "Remoteserviceerror",
		"sub_code": "isv.MOBILE_NUMBER_ILLEGAL",
		"sub_msg": "号码格式错误",
		"request_id": "rxv9u4mi1j9g"
	}
}
*/

func alidayuExample() {
	appkey := "xxx"
	appsecret := "xxx"
	client := DefaultClient(false, false, appkey, appsecret, "")
	smsSendRequest := request.NewAlibabaAliQinFcSmsNumSendRequest("张方", "SMS_56510312")
	smsSendRequest.RecNum = "xxx"
	smsSendRequest.SmsParam = `{"name":"angelina","code":"1234"}`
	res := client.Execute(smsSendRequest)
	fmt.Print(res)
}
