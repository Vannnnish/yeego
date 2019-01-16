/**
 * Created by angelina on 2017/7/29.
 */

package submail

import (
	"encoding/json"
	"errors"
	"testing"
)

func submailSendSmsCode(phoneNum, code string) error {
	config := Config{
		AppId:    "submail.AppId",
		AppKey:   "submail.AppKey",
		SignType: "md5",
	}
	mXSend := CreateMessageXSend(phoneNum, "x25na")
	mXSend.AddVar("code", code)
	result := mXSend.Run(config)
	res := &SubmailResponse{}
	if err := json.Unmarshal([]byte(result), res); err != nil {
		return err
	}
	if res.Status != "success" {
		return errors.New("error")
	}
	return nil
}

func TestMessageXSend_RunVoice(t *testing.T) {
	config := Config{
		AppId:    "xxx",
		AppKey:   "xxx",
		SignType: "md5",
	}
	mXSend := CreateMessageXSend("xxx", "kzKYo1")
	mXSend.AddVar("name", "BTC")
	mXSend.AddVar("price", "8000美元")
	result := mXSend.RunVoice(config)
	res := &SubmailResponse{}
	if err := json.Unmarshal([]byte(result), res); err != nil {
		t.Fatal(err.Error())
	}
	if res.Status != "success" {
		t.Fatal(res.Msg)
	}
	t.Log(res)
}
