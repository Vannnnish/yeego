/**
 * Created by angelina-zf on 17/3/23.
 */

package rsa

import (
	"fmt"
	"testing"
)

const (
	publicKeyPath  string = "client_public.pem"
	privateKeyPath string = "client_private.pem"
)

func TestGenRsaKey(t *testing.T) {
	err := GenRsaKey("", "", 4096)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
}

func TestRsaEncrypt(t *testing.T) {
	rowData := "hello"
	data, err := RsaEncrypt("./client_public.pem", []byte(rowData))
	if err != nil {
		t.Fatal(err)
	}
	decryData, err := RsaDecrypt("./client_private.pem", []byte(data))
	if err != nil {
		t.Fatal(err)
	}
	if rowData != string(decryData) {
		t.Fail()
	}
}
