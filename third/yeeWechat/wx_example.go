/**
 * Created by angelina on 2017/8/2.
 */

package yeeWechat

import (
	"github.com/labstack/echo"
	"github.com/yeeyuntech/yeego"
	"sync"
)

var (
	wxInstance *Wx
	wxOnce     sync.Once
)

func route() {
	// 设置路由
	yeego.Echo.Any(BaseAuthCallbackUrl, baseRedirectPage())
	yeego.Echo.Any(UserInfoAuthCallbackUrl, userInfoRedirectPage())
	yeego.Echo.Any(PaySuccessCallbackUrl, payCallback())
}

// 获取wx实例
func getWx() *Wx {
	wxOnce.Do(func() {
		wxInstance = &Wx{
			AppId:                yeego.Config.GetString("AppId"),
			AppSecret:            yeego.Config.GetString("AppSecret"),
			MchId:                yeego.Config.GetString("MchId"),
			Key:                  yeego.Config.GetString("Key"),
			BaseAuthCallBack:     baseAuthCallBack,
			UserInfoAuthCallBack: userInfoAuthCallBack,
			PaySuccessCallBack:   payCallBack,
		}
	})
	return wxInstance
}

// 基础授权检测
func baseAuth(c echo.Context) {
	getWx().Auth(c, AuthScopeBase, "")
}

// 普通授权回调
func baseRedirectPage() echo.HandlerFunc {
	return getWx().BaseRedirectPage()
}

// 高级授权回调
func userInfoRedirectPage() echo.HandlerFunc {
	return getWx().UserInfoRedirectPage()
}

// 支付成功回调
func payCallback() echo.HandlerFunc {
	return getWx().PaySuccess()
}

// 普通授权回调
func baseAuthCallBack(req *yeego.Request) error {
	res := yeego.NewResponse(req.Context(), req)
	openId := req.SessionGetStr("OpenId")
	if openId == "" {
		panic("invalid openid")
	}
	// find user from db by openid
	user := map[string]string{"OpenId": "xxx"}
	if user["OpenId"] == "" {
		req.SessionSetStr("OpenId", "")
		wxSingle := getWx()
		wxSingle.Auth(res.Context(), AuthScopeUserInfo, "")
		return nil
	}
	if user["OpenId"] != openId {
		panic("openid not match user")
	}
	req.SessionSetStr("OpenId", openId)
	url := req.SessionGetStr("RedirectUrl")
	req.SessionSetStr("RedirectUrl", "")
	if url != "" {
		return res.Redirect(url)
	}
	return res.Redirect("/")
}

// 高级授权回调
func userInfoAuthCallBack(req *yeego.Request) error {
	res := yeego.NewResponse(req.Context(), req)
	openId := req.SessionGetStr("OpenId")
	nickName := req.SessionGetStr("NickName")
	avatar := req.SessionGetStr("Avatar")
	// create user to db
	_ = map[string]string{
		"OpenId":   openId,
		"NickName": nickName,
		"Avatar":   avatar,
	}
	url := req.SessionGetStr("RedirectUrl")
	req.SessionSetStr("RedirectUrl", "")
	if url != "" {
		return res.Redirect(url)
	}
	return res.Redirect("/")
}

// 支付成功回调
func payCallBack(req *yeego.Request, payInfo PaySuccessInfo) {
	return
}

// 检测是否授权
func checkAuth(c echo.Context) (req *yeego.Request, res *yeego.Response) {
	req, res = yeego.NewReqAndRes(c)
	if req.SessionGetStr("OpenId") == "" {
		// 未授权
		redirect := c.Request().URL.String()
		req.SessionSetStr("RedirectUrl", redirect)
		res.SetSession()
		baseAuth(c)
	} else {
		// 已授权
		wx := getWx()
		url := yeego.Config.GetString("SelfSchemeAndDomain") + c.Request().URL.String()
		jsConfig := wx.GetJsConfigSign(url)
		res.Data("JSConfigAppId", jsConfig.AppId)
		res.Data("JSConfigNonceStr", jsConfig.NonceStr)
		res.Data("JSConfigTimeStamp", jsConfig.TimeStamp)
		res.Data("JSConfigSign", jsConfig.Sign)
	}
}

type msgData struct {
	First    msgSubValue `json:"first"`
	Keyword1 msgSubValue `json:"keyword1"` // 捐助金额
	Remark   msgSubValue `json:"remark"`
}

type msgSubValue struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// 推送消息
func sendWxMsg() {
	w := getWx()
	msg := DonateMessage{
		ToUser:     "openId",
		TemplateId: "MsgTypeIdDonate",
		Data: msgData{
			First: msgSubValue{
				Value: "xx",
			},
			Keyword1: msgSubValue{
				Value: "xx",
			},
			Remark: msgSubValue{
				Value: "xx",
			},
		},
	}
	w.SendMessage(msg)
}
