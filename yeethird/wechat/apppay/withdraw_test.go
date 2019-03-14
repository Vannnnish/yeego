/*
@Time : 2019-03-13 23:14 
@Author : vannnnish
@File : withdraw_test
*/

package apppay

import (
	"fmt"
	"testing"
)

func TestWechatTransfer(t *testing.T) {
	Init()
	err := WechatTransfer(appId, mchId, "-sx51CGeiXhfaTh5C8k7JmmJIB0", "NO_CHECK", 30, "ab2323sdfdf", "微信转账", "39.105.27.83", key, "./apiclient_cert.pem", "./apiclient_key.pem")
	fmt.Println("错误:", err)
}
