/**
 * Created by angelina on 2017/8/2.
 * https://mp.weixin.qq.com/advanced/tmplmsg?action=faq&token=634500453&lang=zh_CN
 */

package yeeWechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yeeyuntech/yeego"
	"io/ioutil"
	"net/http"
)

// 默认的捐助模版
type DonateMessage struct {
	ToUser     string      `json:"touser"`
	TemplateId string      `json:"template_id"`
	Url        string      `json:"url"`
	TopColor   string      `json:"topcolor"`
	Data       interface{} `json:"data"`
}

type MessageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int    `json:"msgid"`
}

// SendMessage
// 发送模板消息
func (w *Wx) SendMessage(msg DonateMessage) {
	tmp, _ := json.Marshal(msg)
	token := w.getAccessToken()
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + token
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(tmp)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		yeego.DefaultLogError(fmt.Sprintf("Send Message Error [%s]", err.Error()))
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := MessageResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		yeego.DefaultLogError(fmt.Sprintf("Send Message Error [%s]", err.Error()))
		return
	}
	if response.ErrMsg != "ok" {
		yeego.DefaultLogError(fmt.Sprintf("Send Message Error [%s]", response.ErrMsg))
		return
	}
}
