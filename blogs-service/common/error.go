package common

import (
	"fmt"
)

var (
	// ErrorCode0 正常值
	OkCode = "0"

	// ContextError  上下文异常
	ContextError = "ERR_CONTEXT"

	// SystemErrorCode 系统异常
	SystemErrorCode = "1000"

	// ParamsError 调用参数异常
	ParamsError = NewBizError("1001", "调用参数异常")

	// ErrUnknownError 未知异常
	UnknownError = fmt.Errorf("unknown error")
)

// BizError 业务错误
type BizError struct {
	Code  string
	Msg   string
	Stack string
}

// 实现error接口
func (bize BizError) Error() string {
	return fmt.Sprintf("BizError:%s/%s", bize.Code, bize.Msg)
}

// NewBizError 生成一个BizError
func NewBizError(code string, msg ...string) *BizError {
	var emsg string
	if len(msg) > 0 {
		emsg = msg[0]
	}
	var stack string
	if len(msg) > 1 {
		stack = msg[1]
	}
	return &BizError{code, emsg, stack}
}
