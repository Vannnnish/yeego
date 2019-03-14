/*
@Time : 2019-02-26 00:08 
@Author : vannnnish
@File : withdraw
*/

package apppay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/vannnnish/yeego/yeeparse"
	"github.com/vannnnish/yeego/yeerand"
	"io/ioutil"
	"strings"
)

// 通过微信向店铺付款
var wechatTransferURl = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
// 微信提现(转账)

// 微信转账
// checkName: NO_CHECK 或者 FORCE_CHECK  如果是 FORCE_CHECK 那么 re_user_name 不能为空
func WechatTransfer(mchAppid, mchid, openid, checkName string, amount int, partnerTradeNo, desc, splillCreateIp string, key string, certFile, keyFile string) error {

	noceStr := yeerand.RandString(32)
	wechatTransfer := &WechatWithdrawRequest{
		Nonce_str:        noceStr,
		Partner_trade_no: partnerTradeNo,
		Mch_appid:        mchAppid,
		Mchid:            mchid,
		Openid:           openid,
		Check_name:       checkName,
		// TODO:正式上线再改回来
		Amount:           amount,
		Desc:             desc,
		Spbill_create_ip: splillCreateIp,
	}
	maps := dataToMap(*wechatTransfer)
	wechatTransfer.Sign = Sign(maps, key)
	buf, err := xml.Marshal(wechatTransfer)
	if err != nil {
		return err
	}
	client, err := yeeparse.NewTLSHttpClient(certFile, keyFile)
	if err != nil {
		return err
	}
	tmp := strings.Replace(string(buf), "<WechatOrder>", "<xml>", -1)
	tmp = strings.Replace(tmp, "</WechatOrder>", "</xml>", -1)
	// 转化后的字符串
	body := bytes.NewBuffer([]byte(tmp))
	r, err := client.Post(wechatTransferURl, "text/xml", body)
	if err != nil {
		return err
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	wechatTransferReturn := WechatRefundResult{}
	if err := xml.Unmarshal(response, &wechatTransferReturn); err != nil {
		return err
	}
	if wechatTransferReturn.ResultCode == "SUCCESS" {
		return nil
	} else {
		return errors.New(wechatTransferReturn.ErrCodeDes)
	}
}
