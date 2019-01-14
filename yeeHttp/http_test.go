/*
@Time : 2019-01-14 23:45 
@Author : vannnnish
@File : http_test
*/

package yeeHttp

import (
	"fmt"
	"testing"
)

func TestPayloadRequest(t *testing.T) {
	fmt.Println(PayloadRequest("POST", "http://52.8.148.150:9922/vee/broadcast/payment", `{
  "timestamp": 1547480828021000000,
  "amount": 7700000000,
  "fee": 10000000,
  "feeScale": 100,
  "recipient": "ATxEg46DfzcPmVPpfAUf6GxkrFdABejnmbT",
  "senderPublicKey": "4ZgfGyg4JX5sKLpZ9dXR5yYLMWoarc68gys26WDQs74L",
  "attachment": "",
  "signature": "5K2NqEpdiNPEo5vt8W9MgW1t5mAZdvKftVAT3kVzVpzbqcLTh9bTCRkdhECLDV1fg5CA2SPyz4NvE4L7JZRiMVnp"
}`).Exec().ToString())
}
