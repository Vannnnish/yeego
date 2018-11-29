/**
 * Created by angelina on 2017/7/31.
 */

package yeeAlidayu

type Result struct {
	Code     int
	Msg      string
	Response string
}

type Response struct {
	Res ErrorResponse `json:"error_response"`
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}
