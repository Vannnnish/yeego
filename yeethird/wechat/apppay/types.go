/*
@Time : 2019-01-17 15:46
@Author : vannnnish
@File : types
*/

package apppay

var (
	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?"
)

type Wechat struct {
	AppId     string // 公众账号ID
	AppSecret string // 应用密钥
	MchId     string // 绑定商户ID
	Key       string // 微信商户平台：账户设置-安全设置-API安全-API密钥
}

func NewWechat(appId, appsecret, mchId, key string) Wechat {
	return Wechat{
		AppId:     appId,
		AppSecret: appsecret,
		MchId:     mchId,
		Key:       key,
	}
}

// app支付微信统一下单数据结构
type WechatOrder struct {
	// 必填
	AppId          string `xml:"appid"`            // appId
	MchId          string `xml:"mch_id"`           // 商户号
	NonceStr       string `xml:"nonce_str"`        // 随机字符串
	Sign           string `xml:"sign"`             // 签名
	Body           string `xml:"body"`             // 商品或支付单简要描述
	OutTradeNo     string `xml:"out_trade_no"`     // 商户订单号
	TotalFee       string `xml:"total_fee"`        // 总金额(分)
	SpbillCreateIp string `xml:"spbill_create_ip"` // 终端IP
	NotifyUrl      string `xml:"notify_url"`       // 接收微信支付异步通知回调地址
	TradeType      string `xml:"trade_type"`       // 交易类型 取 APP

	// 非必填
	DeviceInfo string `xml:"device_info"` // 设备号
	SignType   string `xml:"sign_type"`   // 签名类型  默认MD5
	Detail     string `xml:"detail"`      // 商品详情
	Attach     string `xml:"attach"`      // 附加数据
	FeeType    string `xml:"fee_type"`    // 货币类型 CNY
	TimeExpire string `xml:"time_expire"` // 交易结束时间 单号失效时间，  20091227091010   建议:最短失效时间大于1分钟
	GoodsTag   string `xml:"goods_tag"`   // 订单优惠标记
	LimitPay   string `xml:"limit_pay"`   // 指定支付方式，   no_credit--指定不能使用信用卡支付
	Receipt    string `xml:"receipt"`     // 开发票入口开放标识   Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
}

type AppSignStruct struct {
	// 必填
	AppId          string `xml:"appid"`            // appId
	MchId          string `xml:"mch_id"`           // 商户号
	NonceStr       string `xml:"nonce_str"`        // 随机字符串
	Sign           string `xml:"sign"`             // 签名
	Body           string `xml:"body"`             // 商品或支付单简要描述
	OutTradeNo     string `xml:"out_trade_no"`     // 商户订单号
	TotalFee       int    `xml:"total_fee"`        // 总金额(分)
	SpbillCreateIp string `xml:"spbill_create_ip"` // 终端IP
	NotifyUrl      string `xml:"notify_url"`       // 接收微信支付异步通知回调地址
	TradeType      string `xml:"trade_type"`       // 交易类型 取 APP
}

// 统一下单返回数据
type WechatResponse struct {
	// 必定返回
	/*
		COMMENT
			SUCCESS/FAIL此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	*/
	ReturnCode string `xml:"return_code"`
	/*
		COMMENT
			返回信息，如非空，为错误原因:签名失败,参数格式校验错误
	*/
	ReturnMsg string `xml:"return_msg"`
	/*
		COMMENT
			以下信息在return_code为SUCCESS时，必定返回:
	*/
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	Question   string `xml:"-"`
	Solution   string `xml:"-"`
	/*
		COMMENT
			以下字段在return_code 和result_code都为SUCCESS的时候有返回
	*/
	TradeType string `xml:"trade_type"` // JSAPI,NATIVE,APP
	PrepayId  string `xml:"prepay_id"`  // 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
}

// 支付系统错误
/*
	INVALID_REQUEST	参数错误	参数格式有误或者未按规则上传	订单重入时，要求参数值与原请求一致，请确认参数问题
	NOAUTH	商户无此接口权限	商户未开通此接口权限	请商户前往申请此接口权限
	NOTENOUGH	余额不足	用户帐号余额不足	用户帐号余额不足，请用户充值或更换支付卡后再支付
	ORDERPAID	商户订单已支付	商户订单已支付，无需重复操作	商户订单已支付，无需更多操作
	ORDERCLOSED	订单已关闭	当前订单已关闭，无法支付	当前订单已关闭，请重新下单
	SYSTEMERROR	系统错误	系统超时	系统异常，请用相同参数重新调用
	APPID_NOT_EXIST	APPID不存在	参数中缺少APPID	请检查APPID是否正确
	MCHID_NOT_EXIST	MCHID不存在	参数中缺少MCHID	请检查MCHID是否正确
	APPID_MCHID_NOT_MATCH	appid和mch_id不匹配	appid和mch_id不匹配	请确认appid和mch_id是否匹配
	LACK_PARAMS	缺少参数	缺少必要的请求参数	请检查参数是否齐全
	OUT_TRADE_NO_USED	商户订单号重复	同一笔交易不能多次提交	请核实商户订单号是否重复提交
	SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
	XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确
	REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数 	请检查请求参数是否通过post方法提交
	POST_DATA_EMPTY	post数据为空	post数据不能为空	请检查post数据是否为空
	NOT_UTF8	编码格式错误	未使用指定编码格式	请使用NOT_UTF8编码格式
*/
type questionSolution struct {
	question string
	solution string
}

var ErrorInfo = map[string]questionSolution{
	"INVALID_REQUEST":       {"参数格式有误或者未按规则上传", "订单重入时，要求参数值与原请求一致，请确认参数问题"},
	"NOAUTH":                {"商户未开通此接口权限", "请商户前往申请此接口权限"},
	"NOTENOUGH":             {"用户帐号余额不足", "用户帐号余额不足，请用户充值或更换支付卡后再支付"},
	"ORDERPAID":             {"商户订单已支付，无需重复操作", "商户订单已支付，无需更多操作"},
	"ORDERCLOSED":           {"当前订单已关闭，无法支付", "当前订单已关闭，请重新下单"},
	"SYSTEMERROR":           {"系统超时", "系统异常，请用相同参数重新调用"},
	"APPID_NOT_EXIST":       {"参数中缺少APPID", "请检查APPID是否正确"},
	"MCHID_NOT_EXIST":       {"参数中缺少MCHID", "请检查MCHID是否正确"},
	"APPID_MCHID_NOT_MATCH": {"appid和mch_id不匹配", "请确认appid和mch_id是否匹配"},
	"LACK_PARAMS":           {"缺少必要的请求参数", "请检查参数是否齐全"},
	"OUT_TRADE_NO_USED":     {"同一笔交易不能多次提交", "请核实商户订单号是否重复提交"},
	"SIGNERROR":             {"参数签名结果不正确", "请检查签名参数和方法是否都符合签名算法要求"},
	"XML_FORMAT_ERROR":      {"XML格式错误", "请检查XML参数格式是否正确"},
	"REQUIRE_POST_METHOD":   {"未使用post传递参数", "请检查请求参数是否通过post方法提交"},
	"POST_DATA_EMPTY":       {"post数据不能为空", "请检查post数据是否为空"},
	"NOT_UTF8":              {"未使用指定编码格式", "请使用NOT_UTF8编码格式"},
}

func GetError(code string) questionSolution {
	return ErrorInfo[code]
}

// 前台JS下单数据结构
type Order struct {
	OpenId     string
	Body       string // 商品描述
	OutTradeNo string // 商户系统的订单号，与请求一致
	TotalFee   string // 订单总金额，单位为分
	Attach     string // 商家数据包，原样返回
	UserId     string // 用户id
}

// app下单返回数据结构
type PreOrder struct {
	TimeStamp string
	PrepayId  string
}

// 微信退款需要的结构参数
type WechatRefund struct {
	Appid          string `xml:"appid"`          // 公众号ID
	Mch_id         string `xml:"mch_id"`         // 商户号
	Nonce_str      string `xml:"nonce_str"`      // 随机字符串
	Sign           string `xml:"sign"`           // 签名
	Transaction_id string `xml:"transaction_id"` // 微信订单号
	Out_refund_no  string `xml:"out_refund_no"`  // 商户退款号
	Total_fee      string `xml:"total_fee"`      // 订单金额
	Refund_fee     string `xml:"refund_fee"`     // 退款金额
}

// 退款结果返回的字段
type WechatRefundResult struct {
	ResultCode    string `xml:"result_code"`
	ReturnMsg     string `xml:"return_msg"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	TransactionId string `xml:"transaction_id"`
	RefundFee     int    `xml:"refund_fee"`
	RefundId      string `xml:"refund_id"`
}

// 支付成功返回信息
type PaySuccessInfoResponse struct {
	Appid          string `xml:"appid"`
	Bank_type      string `xml:"bank_type"`      // 银行类型
	Cash_fee       string `xml:"cash_fee"`       // 现金支付金额订单现金支付金额
	Fee_type       string `xml:"fee_type"`       // 货币类型
	Is_subscribe   string `xml:"is_subscribe"`   // 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	Mch_id         string `xml:"mch_id"`         //
	Nonce_str      string `xml:"nonce_str"`      //
	Openid         string `xml:"openid"`         //
	Out_trade_no   string `xml:"out_trade_no"`   //
	Result_code    string `xml:"result_code"`    //
	Return_code    string `xml:"return_code"`    //
	Sign           string `xml:"sign"`           // 签名
	Time_end       string `xml:"time_end"`       // 支付完成时间，格式为yyyyMMddHHmmss
	Total_fee      string `xml:"total_fee"`      // 订单总金额
	Trade_type     string `xml:"trade_type"`     // JSAPI  支付类型
	Transaction_id string `xml:"transaction_id"` // 微信支付订单号
	Attach         string `xml:"attach"`         // 返回的ordreId
}

type WechatReturn struct {
	ReturnCode string `xml:"return_code"`
}
type JSConfig struct {
	AppId     string
	NonceStr  string
	TimeStamp string
	Sign      string
}

// 提现接口
type WechatWithdrawRequest struct {
	Mch_appid        string `xml:"mch_appid"`        // 商户号id appid
	Mchid            string `xml:"mchid"`            // 商户号
	Nonce_str        string `xml:"nonce_str"`        // 随机字符串
	Sign             string `xml:"sign"`             // 签名
	Partner_trade_no string `xml:"partner_trade_no"` // 商户订单号
	Openid           string `xml:"openid"`           //
	Check_name       string `xml:"check_name"`       // NO_CHECK：不校验真实姓名  FORCE_CHECK：强校验真实姓名
	Amount           int    `xml:"amount"`           // 金额
	Desc             string `xml:"desc"`             // 企业付款备注
	Spbill_create_ip string `xml:"spbill_create_ip"` // ip地址
}

// 查询支付结果
type QueryWechatRequest struct {
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Out_trade_no string `xml:"out_trade_no"`
	Nonce_str    string `xml:"nonce_str"`
	Sign         string `xml:"sign"`
}

// 支付结果
type QueryWechatResponse struct {
	Return_code  string `xml:"return_code"`
	Return_msg   string `xml:"return_msg"`
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Nonce_str    string `xml:"nonce_str"`
	Result_code  string `xml:"result_code"` // SUCCESS 或者 FAIL
	Err_code     string `xml:"err_code"`
	Err_code_des string `xml:"err_code_des"`
}
