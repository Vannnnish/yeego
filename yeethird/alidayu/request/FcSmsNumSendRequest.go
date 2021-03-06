/**
 * Created by angelina on 2017/7/31.
 */

package request

import "fmt"

// 短信发送
// 向指定手机号码发送模板短信，模板内可设置部分变量。使用前需要在阿里大于管理中心添加短信签名与短信模板。
// alibaba.aliqin.fc.sms.num.send
type AlibabaAliQinFcSmsNumSendRequest struct {
	// 短信类型，传入值请填写normal
	SmsType string
	// 短信签名，传入的短信签名必须是在阿里大于“管理中心-短信签名管理”中的可用签名。
	// 如“阿里大于”已在短信签名管理中通过审核，则可传入”阿里大于“（传参时去掉引号）作为短信签名。短信效果示例：【阿里大于】欢迎使用阿里大于服务。
	SmsFreeSignName string
	// 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。
	// 示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
	// 短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。
	// 群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。示例：18600000000,13911111111,13322222222
	RecNum string
	// 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。
	// 示例：SMS_585014
	SmsTemplateCode string
	// 公共回传参数，在“消息返回”中会透传回该参数；(可选)
	// 举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	Extend string
	// 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。(可选)
	// 示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
	SmsParam string
}

func (r *AlibabaAliQinFcSmsNumSendRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.sms.num.send"
}

func (r *AlibabaAliQinFcSmsNumSendRequest) Check() (err error) {
	defer func() {
		if data := recover(); data != nil {
			err = fmt.Errorf("%v", data)
		}
	}()
	panicIfNil("SmsType", r.SmsType)
	panicIfNil("RecNum", r.RecNum)
	panicIfNil("SmsFreeSignName", r.SmsFreeSignName)
	panicIfNil("SmsTemplateCode", r.SmsTemplateCode)
	err = nil
	return
}

func (r *AlibabaAliQinFcSmsNumSendRequest) GetApiParams() map[string]string {
	params := make(map[string]string)
	params["extend"] = r.Extend
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = r.SmsFreeSignName
	params["sms_param"] = r.SmsParam
	params["rec_num"] = r.RecNum
	params["sms_template_code"] = r.SmsTemplateCode
	return params
}

// NewAlibabaAliQinFcSmsNumSendRequest
// 返回一个sms短信发送请求体
func NewAlibabaAliQinFcSmsNumSendRequest(signName, templateCode string) *AlibabaAliQinFcSmsNumSendRequest {
	return &AlibabaAliQinFcSmsNumSendRequest{
		SmsType:         "noraml",
		SmsFreeSignName: signName,
		SmsTemplateCode: templateCode,
	}
}
