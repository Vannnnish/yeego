/*
@Time : 2019-01-17 21:44 
@Author : vannnnish
@File : wechat_login_test
*/

package applogin

import (
	"encoding/json"
	"fmt"
	"github.com/vannnnish/yeego"
	"testing"
)

var (
	appId     = ""
	appSecret = ""
	mchId     = ""
	key       = ""
	code      = "081mXt880fbwJF1cOA6800ux880mXt8i"
)

func Init() {
	yeego.MustInitConfig("/Users/vannnnish/go_project/project_sh/yeego/conf", "conf")
	appId = yeego.Config.GetString("wechat.AppId")
	appSecret = yeego.Config.GetString("wechat.AppSecret")
	mchId = yeego.Config.GetString("wechat.MchId")
	key = yeego.Config.GetString("wechat.Key")
}

var accessData = `{"access_token":"17_YCm4UU_9j276Z0x1ikElhhWYSEx3pj-tpkFRl7AgsJMf37lhTnDi-BRPd7N9TxGwQc8N5X-k_qd21oBhyMAPKQ2HYrfta9KPi8B7x7zJWOQ","expires_in":7200,"refresh_token":"17_znma7gvAHIM-SBfySyH_oSYIyaQc9BHKrKSYOmHqPSqO5fpvf3yaNN5rb9NrtQi0BCmqbyUurk-GQ7efmk6f23ESOj_aAepKT36up-Uvon8","openid":"o-sx51CI7UcSOMPBgg2Jqq_OsbM4","scope":"snsapi_userinfo","unionid":"os0pQ5nu-cHyGbS0CFE0UAd2Gm_I"}`

func TestWechat(t *testing.T) {
	Init()
	login := NewAppLogin(appId, appSecret)
	ret, err := login.GetAccessToken(code)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("获取到的access_token", ret)
	info, err := ret.GetUserInfo()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("获取到的用户信息", info)
}

func TestLoginReturn_GetUserInfo(t *testing.T) {
	var access LoginReturn
	err := json.Unmarshal([]byte(accessData), &access)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("acc", access)
	userInfo, err := access.GetUserInfo()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("用户信息:", userInfo)

}

func TestAppLogin_RefreshToken(t *testing.T) {
	var access LoginReturn
	err := json.Unmarshal([]byte(accessData), &access)
	if err != nil {
		t.Fatal(err)
	}
	login := NewAppLogin(appId, appSecret)
	_, err = login.RefreshToken(access.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
}
