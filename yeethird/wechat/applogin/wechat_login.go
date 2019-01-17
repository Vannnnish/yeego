/*
@Time : 2019-01-17 20:32 
@Author : vannnnish
@File : types
*/

package applogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vannnnish/yeego/yeehttp"
	"github.com/vannnnish/yeego/yeethird/wechat"
)

/*
	CLUE
		微信登录步骤:
			1. 请求code(这个是调出微信授权窗口后，用户点击确认后，前台用户可以获取到的，然后把这个值传给后台)
			2. 通过code获取access_token
			3. 通过access_token 调用接口

	NOTICE
		注意： 本服务没有保存access_token，如果需要保存请自己实现

*/

const (
	// 获取access_token接口
	AccessTokenApi = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// 刷新token
	RefreshTokenApi = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	// 获取用户信息
	UserInfoApi = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
)

func NewAppLogin(appId, secret string) AppLogin {
	return AppLogin{
		AppId:  appId,
		Secret: secret,
	}
}

// 获取access_token
func (app AppLogin) GetAccessToken(code string) (LoginReturn, error) {
	var ret LoginReturn
	response, err := yeehttp.Get(fmt.Sprintf(AccessTokenApi, app.AppId, app.Secret, code)).
		Exec().ToBytes()
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return ret, err
	}
	if ret.OpenId == "" || ret.UnionId == "" {
		var errData wechat.ErrorData
		err := json.Unmarshal(response, &errData)
		if err != nil {
			return ret, err
		}
		return ret, errors.New(wechat.GetErrorMsg(errData.ErrCode))
	}
	return ret, nil
}

// 刷新token
func (app AppLogin) RefreshToken(refreshToken string) (LoginReturn, error) {
	var ret LoginReturn
	response, err := yeehttp.Get(fmt.Sprintf(RefreshTokenApi, app.AppId, app.Secret)).
		Param("appid", app.AppId).
		Param("grant_type", "refresh_token").
		Param("refresh_token", refreshToken).
		Exec().ToBytes()
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return ret, err
	}
	if ret.OpenId == "" || ret.UnionId == "" {
		var errData wechat.ErrorData
		err := json.Unmarshal(response, &errData)
		if err != nil {
			return ret, err
		}
		return ret, errors.New(wechat.GetErrorMsg(errData.ErrCode))
	}
	return ret, nil
}

// 获取用户信息
func (ret LoginReturn) GetUserInfo() (UserInfo, error) {
	var info UserInfo
	response, err := yeehttp.Get(fmt.Sprintf(UserInfoApi, ret.AccessToken, ret.OpenId)).
		Exec().ToBytes()
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(response, &info)
	if err != nil {
		return info, err
	}
	if info.OpenId == "" || info.UnionId == "" {
		var errData wechat.ErrorData
		err := json.Unmarshal(response, &errData)
		if err != nil {
			return info, err
		}
		return info, errors.New(wechat.GetErrorMsg(errData.ErrCode))
	}
	return info, nil
}
