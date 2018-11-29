/**
 * Created by angelina on 2017/7/31.
 */

package request

type BaseRequest interface {
	// 获取定义的接口名称
	GetApiMethodName() string
	// 入参数检查
	Check() error
	// 获取接口参数
	GetApiParams() map[string]string
}

func panicIfNil(param, data string) {
	if data == "" {
		panic("参数" + param + "不能为空")
	}
}
