package xerr

import "fmt"

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 获取错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 获取错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// Error 实现 error 接口，返回错误信息
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// NewErrCodeMsg (Recommend) 创建自定义错误码和错误信息的错误类型
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errType string) *CodeError {
	return &CodeError{errCode: code[errType], errMsg: errType}
}

// NewErrMsg 返回400，错误信息自定义
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: code[SERVER_COMMON_ERROR], errMsg: errMsg}
}
