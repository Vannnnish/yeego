/**
 * Created by angelina on 2017/7/31.
 */

package yeeAlidayu

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/yeeyuntech/yeego/third/yeeAlidayu/request"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	// 正式的http请求接口
	ProHTTPGateWayURL string = "http://gw.api.taobao.com/router/rest"
	// 沙箱http请求接口
	DebugHTTPGateWayURL string = "http://gw.api.tbsandbox.com/router/rest"
	// 正式的https请求接口
	ProHTTPSGateWayURL string = "https://eco.taobao.com/router/rest"
	// 沙箱的https请求接口
	DebugHTTPSGateWayURL string = "https://gw.api.tbsandbox.com/router/rest"
	// 正确返回会包含的json tag
	successTagJSON string = `"success":true`
)

const (
	ParamsMissError     = -1
	NewRequestError     = -2
	ClientDoError       = -3
	HTTPStatusCodeError = -4
)

// Client
// 是为了可以同时支持多个appkey,可以通过多个client完成请求
type Client struct {
	// 是否是调试环境，调试环境，调用沙箱环境
	Debug bool
	// 是否调用https请求
	HTTPS bool
	// TOP分配给应用的AppKey
	AppKey string
	// AppSecret
	AppSecret string
	// 请求的地址
	GateWayURL string
	// 响应格式。接口不传默认为xml格式，可选值：xml，json。
	// 本sdk默认为json格式
	Format string
	// 签名的摘要算法，可选值为：hmac，md5
	// 默认为md5
	signMethod string
	// API协议版本，
	// 默认：2.0。
	ApiVersion string
}

// DefaultClient
// 生成一个默认的client
func DefaultClient(debug, https bool, appKey, appSecret, gateWayURL string) *Client {
	return &Client{
		Debug:      debug,
		HTTPS:      https,
		AppKey:     appKey,
		AppSecret:  appSecret,
		GateWayURL: gateWayURL,
		Format:     "json",
		signMethod: "md5",
		ApiVersion: "2.0",
	}
}

// Execute
// 执行对应请求
func (c *Client) Execute(request request.BaseRequest) (res Result) {
	// 入参检查
	if err := request.Check(); err != nil {
		res.Code = ParamsMissError
		res.Msg = err.Error()
		return
	}
	// 组装公共参数
	sysParams := make(map[string]string)
	sysParams["method"] = request.GetApiMethodName()
	sysParams["app_key"] = c.AppKey
	sysParams["sign_method"] = c.signMethod
	sysParams["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	sysParams["format"] = c.Format
	sysParams["v"] = c.ApiVersion
	// 业务参数
	apiParams := request.GetApiParams()
	sysParams["sign"] = c.generateMD5Sign(sysParams, apiParams)
	if c.GateWayURL == "" {
		if c.HTTPS {
			c.GateWayURL = ProHTTPSGateWayURL
		} else {
			c.GateWayURL = ProHTTPGateWayURL
		}
	}
	requestURL := c.GateWayURL + "?"
	for k, v := range sysParams {
		requestURL += k + "=" + url.QueryEscape(v) + "&"
	}
	requestURL = requestURL[:len(requestURL)-1]
	v := url.Values{}
	var keys []string
	for k := range apiParams {
		keys = append(keys, k)
	}
	for _, k := range keys {
		v.Set(k, apiParams[k])
	}
	return c.curl(requestURL, bytes.NewBuffer([]byte(v.Encode())))
}

func (c *Client) curl(requestURL string, body io.Reader) (res Result) {
	client := &http.Client{}
	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", requestURL, body)
	if err != nil {
		res.Code = NewRequestError
		res.Msg = err.Error()
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		res.Code = ClientDoError
		res.Msg = err.Error()
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response := string(data)
	if resp.StatusCode != 200 {
		res.Code = HTTPStatusCodeError
		res.Msg = "http状态码不是200"
		res.Response = response
		return
	}
	// 解析response
	r := &Response{}
	if strings.Contains(response, "error_response") {
		json.Unmarshal([]byte(response), r)
		res.Code = r.Res.Code
		res.Msg = r.Res.Msg
		return
	}
	res.Code = 0
	if strings.Contains(response, successTagJSON) {
		res.Msg = "success"
	}
	res.Response = response
	return
}

func (c *Client) generateMD5Sign(sysParams, apiParams map[string]string) string {
	for k, v := range sysParams {
		apiParams[k] = v
	}
	var keys []string
	for k := range apiParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signString := c.AppSecret
	for _, k := range keys {
		signString += k + apiParams[k]
	}
	signString += c.AppSecret
	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	return sign
}
