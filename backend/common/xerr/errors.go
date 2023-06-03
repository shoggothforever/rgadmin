package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code uint32
	Msg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Msg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{Code: errCode, Msg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{Code: SERVER_COMMON_ERROR, Msg: errMsg}
}
