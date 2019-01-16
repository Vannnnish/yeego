package jpush

import (
	"encoding/base64"
	"encoding/json"
	"github.com/vannnnish/yeego/third/jpush/common"
)

const (
	SUCCESS_FLAG           = "msg_id"
	HOST_NAME_SSL          = "https://api.jpush.cn/v3/push"
	BASE64_TABLE           = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	VALIDATE_HOST_NAME_SSL = "https://api.jpush.cn/v3/push/validate"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *PushClient {
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{secret, appKey, auth, HOST_NAME_SSL}
	return pusher
}

func (pushClient *PushClient) Send(builder interface{}) (*common.RetData, error) {
	content, err := json.Marshal(builder)
	if err != nil {
		return nil, err
	}
	return pushClient.SendPushBytes(content)
}

func (pushClient *PushClient) SendPushString(content string) (ret *common.RetData, err error) {
	ret, err = common.SendPostString(pushClient.BaseUrl, content, pushClient.AuthCode)
	return
}

func (pushClient *PushClient) SendPushBytes(content []byte) (ret *common.RetData, err error) {
	ret, err = common.SendPostBytes(pushClient.BaseUrl, content, pushClient.AuthCode)
	return
}

func (pushClient *PushClient) SendValidate(builder interface{}) (*common.RetData, error) {
	content, err := json.Marshal(builder)
	if err != nil {
		return nil, err
	}
	return pushClient.SendValidatePushBytes(content)
}

func (pushClient *PushClient) SendValidatePushBytes(content []byte) (ret *common.RetData, err error) {
	ret, err = common.SendPostBytes(VALIDATE_HOST_NAME_SSL, content, pushClient.AuthCode)
	return
}
