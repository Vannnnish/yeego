/**
 * Created by angelina on 2017/5/27.
 */

package submail

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/vannnnish/yeego/yeehttp"
	"github.com/vannnnish/yeego/yeestrconv"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Config struct {
	AppId    string
	AppKey   string
	SignType string
}

// submail短信验证码返回参数
type SubmailResponse struct {
	Status     string  `json:"status"`                // 状态
	SendId     string  `json:"send_id,omitempty"`     // 此处发送的id
	Fee        float64 `json:"fee,omitempty"`         //
	SmsCredits string  `json:"sms_credits,omitempty"` //
	Code       int     `json:"code,omitempty"`        //
	Msg        string  `json:"msg,omitempty"`         // 错误原因?
}

func httpPost(queryUrl string, postData map[string]string) string {
	data, err := json.Marshal(postData)
	if err != nil {
		return err.Error()
	}
	body := bytes.NewBuffer([]byte(data))
	retStr, err := http.Post(queryUrl, "application/json;charset=utf-8", body)
	if err != nil {
		return err.Error()
	}
	result, err := ioutil.ReadAll(retStr.Body)
	retStr.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func getTimeStamp() string {
	resp, err := yeehttp.Get("https://api.submail.cn/service/timestamp.json").Exec().ToBytes()
	if err != nil {
		return err.Error()
	}
	var dict map[string]interface{}
	err = json.Unmarshal(resp, &dict)
	if err != nil {
		return err.Error()
	}
	return yeestrconv.FormatInt(int(dict["timestamp"].(float64)))
}

func createSignature(request map[string]string, config Config) string {
	appKey := config.AppKey
	appId := config.AppId
	signType := config.SignType
	request["sign_type"] = signType
	keys := make([]string, 0, 32)
	for key := range request {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str_list := make([]string, 0, 32)
	for _, key := range keys {
		str_list = append(str_list, fmt.Sprintf("%s=%s", key, request[key]))
	}
	sigstr := strings.Join(str_list, "&")
	sigstr = fmt.Sprintf("%s%s%s%s%s", appId, appKey, sigstr, appId, appKey)
	if signType == "normal" {
		return appKey
	} else if signType == "md5" {
		myMd5 := md5.New()
		io.WriteString(myMd5, sigstr)
		return fmt.Sprintf("%x", myMd5.Sum(nil))
	} else {
		mySha1 := sha1.New()
		io.WriteString(mySha1, sigstr)
		return fmt.Sprintf("%x", mySha1.Sum(nil))
	}
}

type MessageXSend struct {
	to      string
	project string
	vars    map[string]string
}

func CreateMessageXSend(to, project string) *MessageXSend {
	messageXSend := new(MessageXSend)
	messageXSend.to = to
	messageXSend.project = project
	messageXSend.vars = make(map[string]string)
	return messageXSend
}

func (messageXSend *MessageXSend) AddVar(key string, value string) {
	messageXSend.vars[key] = value
}

func (messageXSend *MessageXSend) buildRequest() map[string]string {
	request := make(map[string]string)
	request["to"] = messageXSend.to
	request["project"] = messageXSend.project
	if len(messageXSend.vars) != 0 {
		data, err := json.Marshal(messageXSend.vars)
		if err == nil {
			request["vars"] = string(data)
		}
	}
	return request
}

func (messageXSend *MessageXSend) Run(config Config) string {
	request := messageXSend.buildRequest()
	url := "https://api.mysubmail.com/message/xsend"
	request["appid"] = config.AppId
	request["timestamp"] = getTimeStamp()
	request["signature"] = createSignature(request, config)
	return httpPost(url, request)
}

func (messageXSend *MessageXSend) RunInternational(config Config) string {
	request := messageXSend.buildRequest()
	url := "https://api.mysubmail.com/internationalsms/xsend"
	request["appid"] = config.AppId
	request["timestamp"] = getTimeStamp()
	request["signature"] = createSignature(request, config)
	return httpPost(url, request)
}

func (messageXSend *MessageXSend) RunVoice(config Config) string {
	request := messageXSend.buildRequest()
	url := "https://api.mysubmail.com/voice/xsend"
	request["appid"] = config.AppId
	request["timestamp"] = getTimeStamp()
	request["signature"] = createSignature(request, config)
	return httpPost(url, request)
}
