/**
 * Created by angelina on 2017/8/2.
 * https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842
 */

package yeeWechat

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeStrconv"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 认证类型
type AuthScope string

const (
	// 基础认证
	AuthScopeBase AuthScope = "snsapi_base"
	// 高级认证
	AuthScopeUserInfo AuthScope = "snsapi_userinfo"
	// 基础认证回调地址
	BaseAuthCallbackUrl = "/Wx/BaseRedirectPage"
	// 高级认证回调地址
	UserInfoAuthCallbackUrl = "/Wx/UserInfoRedirectPage"
)

type AuthUrl struct {
	AppId       string
	RedirectURL string
	Scope       string // snsapi_base 或者 snsapi_userinfo
	State       string
}

func getAuthUrl(authUrl AuthUrl) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize" +
		"?appid=" + url.QueryEscape(authUrl.AppId) +
		"&redirect_uri=" + url.QueryEscape(authUrl.RedirectURL) +
		"&response_type=code&scope=" + url.QueryEscape(authUrl.Scope) +
		"&state=" + url.QueryEscape(authUrl.State) +
		"#wechat_redirect"
}

// Auth
// 跳转微信认证页面
func (w *Wx) Auth(c echo.Context, scope AuthScope, state string) {
	_, res := yeego.NewReqAndRes(c)
	redirect := yeego.Config.GetString("SelfSchemeAndDomain") + BaseAuthCallbackUrl
	if scope == AuthScopeUserInfo {
		redirect = yeego.Config.GetString("SelfSchemeAndDomain") + UserInfoAuthCallbackUrl
	}
	authUrl := getAuthUrl(AuthUrl{
		AppId:       w.AppId,
		RedirectURL: redirect,
		Scope:       string(scope),
		State:       state,
	})
	res.Redirect(authUrl)
}

// 基础信息跳转页面
// 获取openid以及token 然后跳转
func (w *Wx) BaseRedirectPage() echo.HandlerFunc {
	f := func(c echo.Context) error {
		req, _ := yeego.NewReqAndRes(c)
		baseInfo := w.GetBaseInfo(req)
		req.SessionSetStr("OpenId", baseInfo.OpenId)
		return w.BaseAuthCallBack(req)
	}
	return f
}

// 基础信息
type BaseInfo struct {
	OpenId string `json:"openid"`
	Token  string `json:"access_token"`
}

// 获取用户基础信息
func (w *Wx) GetBaseInfo(req *yeego.Request) BaseInfo {
	code := req.Param("code").GetString()
	if code == "" {
		yeego.DefaultLogError("Wx GetBaseInfo code empty")
		panic("code不能为空")
	}
	requestUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + w.AppId +
		"&secret=" + w.AppSecret + "&code=" + code + "&grant_type=authorization_code"
	body := w.Request(requestUrl)
	baseInfo := BaseInfo{}
	err := json.Unmarshal(body, &baseInfo)
	if err != nil {
		yeego.DefaultLogError(fmt.Sprintf("Wx GetBaseInfo Error :[%s]", err.Error()))
		panic(err)
	}
	if baseInfo.OpenId == "" {
		yeego.DefaultLogError("Wx GetBaseInfo Get OpenId failed")
		panic("Wx: Get OpenId failed")
	}
	return baseInfo
}

// 用户详细信息跳转页面
// 获取用户详细信息
func (w *Wx) UserInfoRedirectPage() echo.HandlerFunc {
	f := func(c echo.Context) error {
		req, _ := yeego.NewReqAndRes(c)
		baseInfo := w.GetBaseInfo(req)
		userInfo := w.GetUserInfo(baseInfo)
		req.SessionSetStr("OpenId", userInfo.OpenId)
		req.SessionSetStr("NickName", userInfo.NickName)
		req.SessionSetStr("Avatar", userInfo.Avatar)
		req.SessionSetStr("Sex", yeeStrconv.FormatInt(userInfo.Sex))
		return w.UserInfoAuthCallBack(req)
	}
	return f
}

// UserInfo
// 用户详细信息
type UserInfo struct {
	OpenId   string `json:"openid"`
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	Avatar   string `json:"headimgurl"`
}

// 获取用户信息
func (w *Wx) GetUserInfo(baseInfo BaseInfo) UserInfo {
	requestUrl := "https://api.weixin.qq.com/sns/userinfo?access_token=" +
		baseInfo.Token + "&openid=" + baseInfo.OpenId + "&lang=zh_CN"
	body := w.Request(requestUrl)
	userInfo := UserInfo{}
	err := json.Unmarshal(body, &userInfo)
	if err != nil {
		yeego.DefaultLogError(fmt.Sprintf("get baseinfo error :[%s]", err.Error()))
		panic(err)
	}
	if baseInfo.OpenId == "" {
		yeego.DefaultLogError("Wechat: Get user info failed, invalid openid")
		panic("Wechat: Get user info failed, invalid openid")
	}
	return userInfo
}

// 请求数据
func (w *Wx) Request(requestUrl string) []byte {
	client := http.Client{}
	formResponse, err := client.PostForm(requestUrl, url.Values{})
	if err != nil {
		panic(err)
	}
	defer formResponse.Body.Close()
	body, err := ioutil.ReadAll(formResponse.Body)
	if err != nil {
		panic(err)
	}
	return body
}
