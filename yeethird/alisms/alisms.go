/*
@Time : 2019-02-24 05:03 
@Author : vannnnish
@File : alisms
*/

package alisms

import (
	"errors"
	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/satori/go.uuid"
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeejson"
)

var (
	SmsCodeTypeLogin     = "login"
	SmsCodeTypeResetPwd  = "reset_pwd"
	SmsCodeTypeBindPhone = "bind_phone"

	// 国内短信模板
	chinaLoginTpl = map[string]string{
		EnvZhCn:   "SMS_144450880",
		EnvZhTw:   "SMS_144455971",
		EnvEnUs:   "SMS_144456182",
		EnvKoKr:   "SMS_144456176",
		EnvJaJp:   "SMS_144456160",
		"default": "SMS_144456182",
	}
	chinaResetPwdTpl = map[string]string{
		EnvZhCn:   "SMS_152541104",
		EnvEnUs:   "SMS_152546096",
		"default": "SMS_152546096",
	}
	chinaBindPhoneTpl = map[string]string{
		EnvZhCn:   "SMS_152541294",
		EnvEnUs:   "SMS_152546133",
		"default": "SMS_152546133",
	}
	// 国际短信模板
	internationalLoginTpl = map[string]string{
		EnvZhCn:   "SMS_145592870",
		EnvZhTw:   "SMS_145597874",
		EnvEnUs:   "SMS_145597868",
		EnvKoKr:   "SMS_145597873",
		EnvJaJp:   "SMS_145592869",
		"default": "SMS_145597868",
	}
	internationalResetPwdTpl = map[string]string{
		EnvZhCn:   "SMS_152541263",
		EnvZhTw:   "SMS_152546105",
		EnvEnUs:   "SMS_152541267",
		"default": "SMS_152541267",
	}
	internationalBindPhoneTpl = map[string]string{
		EnvZhCn:   "SMS_152541300",
		EnvZhTw:   "SMS_152541303",
		EnvEnUs:   "SMS_152546141",
		"default": "SMS_152546141",
	}
	// 短信签名(不能改，这个是要审核的)
	bullseyeSignName = "bullseye"
	niuyanSignName   = "牛眼行情"
	niuyanEnSignName = "NIUYAN"
)

// 发送手机登录验证码
// 发送国际/港澳台消息时，接收号码格式为00+国际区号+号码，
func SendSmsCode(isNiuyan bool, t, lan, areaCode, phoneNum, code string) error {
	dysms.HTTPDebugEnable = false
	accessId := yeego.Config.GetString("alisms.AccessKey")
	accessKey := yeego.Config.GetString("alisms.AccessSecret")
	dysms.SetACLClient(accessId, accessKey)
	var tplCode string
	switch t {
	case SmsCodeTypeLogin:
		tplCode = getLoginTplCode(lan, areaCode)
	case SmsCodeTypeResetPwd:
		tplCode = getResetPwdTplCode(lan, areaCode)
	case SmsCodeTypeBindPhone:
		tplCode = getBindPhoneTplCode(lan, areaCode)
	default:
		return errors.New("不支持的type")
	}
	param := map[string]string{
		"code": code,
	}
	paramBytes, _ := yeejson.Json().Marshal(param)
	if areaCode != "86" {
		phoneNum = "00" + areaCode + phoneNum
	}
	signName := bullseyeSignName
	if isNiuyan {
		if areaCode == "86" {
			signName = niuyanSignName
		} else {
			signName = niuyanEnSignName
		}
	}
	resp, err := dysms.SendSms(uuid.NewV1().String(), phoneNum, signName, tplCode, string(paramBytes)).DoActionWithException()
	if err != nil {
		return err
	}
	if resp.GetCode() != "OK" {
		return errors.New(resp.GetMessage())
	}
	return nil
}

func getLoginTplCode(lan, areaCode string) string {
	var tpl map[string]string
	if areaCode == "86" {
		tpl = chinaLoginTpl
	} else {
		tpl = internationalLoginTpl
	}
	code, ok := tpl[lan]
	if ok {
		return code
	}
	return tpl["default"]
}

func getResetPwdTplCode(lan, areaCode string) string {
	var tpl map[string]string
	if areaCode == "86" {
		tpl = chinaResetPwdTpl
	} else {
		tpl = internationalResetPwdTpl
	}
	code, ok := tpl[lan]
	if ok {
		return code
	}
	return tpl["default"]
}

func getBindPhoneTplCode(lan, areaCode string) string {
	var tpl map[string]string
	if areaCode == "86" {
		tpl = chinaBindPhoneTpl
	} else {
		tpl = internationalBindPhoneTpl
	}
	code, ok := tpl[lan]
	if ok {
		return code
	}
	return tpl["default"]
}
