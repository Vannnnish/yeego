/*
@Time : 2019-01-17 20:58 
@Author : vannnnish
@File : types
*/

package applogin

type AppLogin struct {
	AppId  string
	Secret string
}

// 返回的数据结构
type LoginReturn struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"string"`
}

//  用户信息
type UserInfo struct {
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"` // 1为男性 2为女性
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	HeadImgUrl string `json:"headimgurl"`
	UnionId    string `json:"unionid"`
}
