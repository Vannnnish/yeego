/**
 * Created by angelina on 2017/5/27.
 */

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// GenRsaKey
// 生成rsa公钥以及私钥
// @params publicKeyPath 公钥地址
// @params privateKeyPath 私钥地址
// @params bits 密钥长rsa_test.go度
func GenRsaKey(publicKeyPath, privateKeyPath string, bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(privateKeyPath + "private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(publicKeyPath + "public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	return err
}

// RsaEncrypt
// rsa加密
// @params publicKeyPath 公钥文件位置
// @params origData 需要加密的数据
// @return 返回的数据是经过base64编码的
func RsaEncrypt(publicKeyPath string, origData []byte) (string, error) {
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", fmt.Errorf("read public key yeefile: %s", err)
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	returnData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	return base64.StdEncoding.EncodeToString(returnData), err
}

// RsaDecrypt
// rsa解密
// 需要对传入的数据进行base64解码
// @params privateKeyPath 私钥文件位置
// @params ciphertext 需要解密的数据
func RsaDecrypt(privateKeyPath string, ciphertext []byte) ([]byte, error) {
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read private key yeefile: %s", err)
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	base64Text, _ := base64.StdEncoding.DecodeString(string(ciphertext))
	return rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(base64Text))
}
