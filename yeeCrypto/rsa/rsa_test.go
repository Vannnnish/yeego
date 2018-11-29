/**
 * Created by angelina-zf on 17/3/23.
 */

package rsa

import (
	"fmt"
	"testing"
)

const (
	publicKeyPath  string = "public.pem"
	privateKeyPath string = "private.pem"
)

func TestGenRsaKey(t *testing.T) {
	err := GenRsaKey("", "", 4096)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
}
