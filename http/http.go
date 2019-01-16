/**
 * Created by angelina on 2017/4/25.
 */

package http

import (
	"crypto/tls"
	"github.com/vannnnish/yeego/file"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// RequestWrapper
// 对http请求的封装
type RequestWrapper struct {
	client   *http.Client
	request  *http.Request
	response *http.Response
	err      error
}

func NewNewRequestWithHttpReq(req *http.Request) *RequestWrapper {
	return &RequestWrapper{request: req}
}

// NewRequest
// 初始化请求
func NewRequest(method, url string) *RequestWrapper {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	return &RequestWrapper{request: req}
}
func NewRequestWithBody(method, url string, body io.Reader) *RequestWrapper {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	return &RequestWrapper{request: req}
}

// Get
// 初始化get请求
func Get(url string) *RequestWrapper {
	return NewRequest("GET", url)
}

// Post
// 初始化post请求
func Post(url string) *RequestWrapper {
	return NewRequest("POST", url)
}

// 发起Payload请求
func PayloadRequest(method, url, data string) *RequestWrapper {
	payload := strings.NewReader(data)
	return NewRequestWithBody(method, url, payload)
}

func (r *RequestWrapper) Query(query string) *RequestWrapper {
	r.request.URL.RawQuery = query
	return r
}

// Param
// 添加url参数
func (r *RequestWrapper) Param(key, value string) *RequestWrapper {
	if r.request.Method == "GET" {
		query := r.request.URL.Query()
		query.Add(key, value)
		return r.Query(query.Encode())
	}
	if r.request.Method == "POST" {
		if r.request.PostForm == nil {
			r.request.PostForm = url.Values{}
		}
		r.request.PostForm.Add(key, value)
	}
	return r
}

// SetHeader
// 设置header
func (r *RequestWrapper) SetHeader(key, value string) *RequestWrapper {
	r.request.Header.Set(key, value)
	return r
}

func (r *RequestWrapper) AddHeader(key, value string) *RequestWrapper {
	r.request.Header.Add(key, value)
	return r
}

// UseClient
// 设置client
func (r *RequestWrapper) UseClient(client *http.Client) *RequestWrapper {
	if client != nil {
		r.client = client
	}
	return r
}

// Json
// 设置post请求的参数 body主体
// `{"greeting":"hello world"}`
func (r *RequestWrapper) Json(data string) *RequestWrapper {
	reader := strings.NewReader(data)
	r.request.Body = ioutil.NopCloser(reader)
	r.request.ContentLength = int64(reader.Len())
	r.SetHeader("Content-Type", "application/json")
	return r
}

// Exec
// 执行http请求
func (r *RequestWrapper) Exec() *RequestWrapper {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := http.DefaultClient
	//client.Transport = tr
	if r.client != nil {
		client = r.client
	}
	client.Timeout = time.Second * 60
	if r.request.Method == http.MethodPost && r.request.PostForm != nil {
		body := io.Reader(strings.NewReader(strings.TrimSpace(r.request.PostForm.Encode())))
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = ioutil.NopCloser(body)
		}
		r.request.Body = rc
		r.request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	r.response, r.err = client.Do(r.request)
	return r
}

// ToBytes
// 结果输出为[]byte
func (r *RequestWrapper) ToBytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.response.Body.Close()
	return ioutil.ReadAll(r.response.Body)
}

// ToString
// 结果输出为string
func (r *RequestWrapper) ToString() (string, error) {
	data, err := r.ToBytes()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Pipe
// 结果输出为io.Writer
// 返回length 以及error
func (r *RequestWrapper) Pipe(w io.Writer) (written int64, err error) {
	if r.err != nil {
		return 0, r.err
	}
	defer r.response.Body.Close()
	written, err = io.Copy(w, r.response.Body)
	return
}

// ToFile
// 结果输出到文件
func (r *RequestWrapper) ToFile(filename string) (size int64, err error) {
	if err := file.MkdirForFile(filename); err != nil {
		return 0, err
	}
	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	size, err = r.Pipe(file)
	return
}

// 返回response
func (r *RequestWrapper) Response() (*http.Response, error) {
	return r.response, r.err
}

// Download
// 下载文件
func Download(url, filename string) (size int64, err error) {
	size, err = Get(url).Exec().ToFile(filename)
	return
}

// 忽略tls认证
func (r *RequestWrapper) SkipVerify() *RequestWrapper {
	if r.client == nil {
		client := http.DefaultClient
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
		r.client = client
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		r.client.Transport = tr
	}
	return r
}
