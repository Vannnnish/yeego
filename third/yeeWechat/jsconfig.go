/**
 * Created by angelina on 2017/8/2.
 */

package yeeWechat

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeHttp"
	"github.com/yeeyuntech/yeego/yeeRand"
	"strconv"
	"sync"
	"time"
)

type Ticket struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Ticket  string `json:"ticket"`
	Expire  int    `json:"expires_in"`
}

type JSConfig struct {
	AppId     string
	NonceStr  string
	TimeStamp string
	Sign      string
}

// 签名使用的nonceStr，timestamp和JS配置中必须相同
// JS SDK config
func (w *Wx) GetJsConfigSign(currentUrl string) JSConfig {
	token := w.getAccessToken()
	content, err := yeeHttp.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + token + "&type=jsapi").
		Exec().ToBytes()
	if err != nil {
		yeego.DefaultLogError(err)
	}
	ticket := Ticket{}
	json.Unmarshal(content, &ticket)
	nonceStr := yeeRand.RandString(16)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := "jsapi_ticket=" + ticket.Ticket + "&noncestr=" + nonceStr + "&timestamp=" + timeStamp + "&url=" + currentUrl
	s := sha1.New()
	s.Write([]byte(str))
	return JSConfig{
		AppId:     w.AppId,
		NonceStr:  nonceStr,
		TimeStamp: timeStamp,
		Sign:      hex.EncodeToString(s.Sum(nil)),
	}
}

type TokenInfo struct {
	Token   string `json:"access_token"`
	Expires int    `json:"expires_in"`
}

type tokenCache struct {
	sync.RWMutex
	Token   string
	GetTime time.Time
}

var token tokenCache

func (w *Wx) getAccessToken() string {
	token.Lock()
	defer token.Unlock()
	// 缓存
	if token.Token != "" && time.Now().Unix()-token.GetTime.Unix() < 7100 {
		return token.Token
	}
	content, _ := yeeHttp.Get("https://api.weixin.qq.com/cgi-bin/token?").
		Param("grant_type", "client_credential").
		Param("appid", w.AppId).
		Param("secret", w.AppSecret).
		Exec().ToBytes()
	response := TokenInfo{}
	json.Unmarshal(content, &response)
	if response.Token == "" {
		panic("获取access_token失败")
	}
	token.Token = response.Token
	token.GetTime = time.Now()
	return response.Token
}
