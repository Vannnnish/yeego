/*
@Time : 2019-02-24 05:03 
@Author : vannnnish
@File : alisms_test
*/

package alisms

import (
	"github.com/vannnnish/yeego"
	"testing"
)

func TestSendLoginSmsCode(t *testing.T) {
	yeego.MustInitConfig("/Users/vannnnish/go_project/project_sh/yeego/conf", "conf")
	err := SendSmsCode(false, "zh-cn", SmsCodeTypeLogin, "86", "17318070950", "1234321")
	if err != nil {
		t.Fatal(err.Error())
	}
}
