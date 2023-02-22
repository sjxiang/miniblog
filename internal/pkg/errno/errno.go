package errno

import "fmt"

type ErrNo struct {
	Status  int    // HTTP 状态码
	BizCode string // 业务错误码  （为啥不是打日志呢？感觉挺扯淡的）
	Message string // 可直接暴露给用户的错误信息   
}

// 实现 error 接口中的 `Error` 方法
func (e *ErrNo) Error() string {
	return e.Message
}

func (e *ErrNo) WithMessage(format string, args ...interface{}) *ErrNo {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

// 尝试从 err 中解析出业务错误码和错误信息
func Decode(err error) (int, string, string) {
	if err == nil {
		return Ok.Status, Ok.BizCode, Ok.Message
	}

	switch typed := err.(type) {
	case *ErrNo:
		return typed.Status, typed.BizCode, typed.Message
	default:
	}

	// 默认返回位置错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.Status, InternalServerError.BizCode, err.Error()
}
