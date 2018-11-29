/**
 * Created by angelina on 2017/5/4.
 */

package yeeJpush

import (
	"errors"
	"fmt"
)

type Client struct {
	c            *PushClient
	IsIosProduct bool
	name         string
}

type NewClientRequest struct {
	Name         string //这个客户端的名字
	AppKey       string //
	Secret       string //
	IsIosProduct bool   //如果是否false表示向测试设备推送,如果是true表示向正式设备推送,后台的那个开发与正式似乎没有作用.
}

func NewClient(req NewClientRequest) *Client {
	return &Client{
		c:            NewPushClient(req.Secret, req.AppKey),
		IsIosProduct: req.IsIosProduct,
		name:         req.Name,
	}
}

var NotFoundUser error = errors.New("[Jpush] user not exist")

func (c *Client) PushToSome(alias []string, subTitle, content string, extras map[string]interface{}) (err error) {
	nb := NewNoticeBuilder()
	nb.SetPlatform(AllPlatform())
	au := &Audience{}
	au.SetAlias(alias)
	nb.SetAudience(au)

	notice := NewNoticeAndroid()
	notice.Alert = content
	notice.Extras = extras
	nb.SetNotice(notice)

	body := map[string]string{
		"title": subTitle,
		"body":  content,
	}

	iosNotice := NewNoticeIos()
	iosNotice.Sound = "default"
	iosNotice.Badge = "1"
	iosNotice.Alert = body
	iosNotice.Extras = extras
	nb.SetNotice(iosNotice)

	op := NewOptions()
	op.SetApns_production(c.IsIosProduct)
	nb.SetOptions(op)
	ret, err := c.c.Send(nb)
	if err != nil {
		return err
	}
	if ret.Error.Code == 0 {
		return nil
	}
	if ret.Error.Code == 1011 {
		return NotFoundUser
	}
	return fmt.Errorf("code:%d err: %s", ret.Error.Code, ret.Error.Message)
}

func (c *Client) ValidatePushToSome(alias []string, content string) (err error) {
	nb := NewNoticeBuilder()
	nb.SetPlatform(AllPlatform())
	au := &Audience{}
	au.SetAlias(alias)
	nb.SetAudience(au)

	notice := NewNotice()
	notice.Alert = content
	nb.SetNotice(notice)

	iosNotice := NewNoticeIos()
	iosNotice.Sound = "default"
	iosNotice.Badge = "1"
	iosNotice.Alert = content
	nb.SetNotice(iosNotice)

	op := NewOptions()
	op.SetApns_production(c.IsIosProduct)
	nb.SetOptions(op)
	ret, err := c.c.SendValidate(nb)
	if err != nil {
		return err
	}
	if ret.Error.Code == 0 {
		return nil
	}
	if ret.Error.Code == 1011 {
		return NotFoundUser
	}
	return fmt.Errorf("code:%d err: %s", ret.Error.Code, ret.Error.Message)
}

func (c *Client) PushToOne(alias, subTitle, content string, extras map[string]interface{}) (err error) {
	nb := NewNoticeBuilder()
	nb.SetPlatform(AllPlatform())
	au := &Audience{}
	au.SetAlias([]string{alias})
	nb.SetAudience(au)

	notice := NewNoticeAndroid()
	notice.Alert = content
	notice.Extras = extras
	nb.SetNotice(notice)

	body := map[string]string{
		"title": subTitle,
		"body":  content,
	}

	iosNotice := NewNoticeIos()
	iosNotice.Sound = "default"
	iosNotice.Badge = "1"
	iosNotice.Alert = body
	iosNotice.Extras = extras
	nb.SetNotice(iosNotice)

	op := NewOptions()
	op.SetApns_production(c.IsIosProduct)
	nb.SetOptions(op)
	ret, err := c.c.Send(nb)
	if err != nil {
		return err
	}
	if ret.Error.Code == 0 {
		return nil
	}
	if ret.Error.Code == 1011 {
		return NotFoundUser
	}
	return fmt.Errorf("code:%d err: %s", ret.Error.Code, ret.Error.Message)
}

func (c *Client) ValidatePushToOne(alias, content string) (err error) {
	nb := NewNoticeBuilder()
	nb.SetPlatform(AllPlatform())
	au := &Audience{}
	au.SetAlias([]string{alias})
	nb.SetAudience(au)

	notice := NewNotice()
	notice.Alert = content
	nb.SetNotice(notice)

	iosNotice := NewNoticeIos()
	iosNotice.Sound = "default"
	iosNotice.Badge = "1"
	iosNotice.Alert = content
	nb.SetNotice(iosNotice)

	op := NewOptions()
	op.SetApns_production(c.IsIosProduct)
	nb.SetOptions(op)
	ret, err := c.c.SendValidate(nb)
	if err != nil {
		return err
	}
	if ret.Error.Code == 0 {
		return nil
	}
	if ret.Error.Code == 1011 {
		return NotFoundUser
	}
	return fmt.Errorf("code:%d err: %s", ret.Error.Code, ret.Error.Message)
}

func (c *Client) PushToAll(content string) (err error) {

	nb := NewNoticeBuilder()
	nb.SetPlatform(AllPlatform())
	nb.SetAudience(AllAudience())

	notice := NewNoticeAndroid()
	notice.Alert = content
	nb.SetNotice(notice)

	iosNotice := NewNoticeIos()
	iosNotice.Sound = "default"
	iosNotice.Badge = "1"
	iosNotice.Alert = content
	nb.SetNotice(iosNotice)

	op := NewOptions()
	op.SetApns_production(c.IsIosProduct)
	op.SetBigPushDuration(60) //过快的进行全局推送,会导致系统其他地方压力太大而挂掉.先设置成60分钟.
	nb.SetOptions(op)
	ret, err := c.c.Send(nb)
	if err != nil {
		return err
	}
	if ret.Error.Code == 0 {
		return nil
	}
	if ret.Error.Code == 1011 {
		return NotFoundUser
	}
	return fmt.Errorf("code:%d err: %s", ret.Error.Code, ret.Error.Message)
}
