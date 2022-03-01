package errors

import (
	"fmt"
)

type Error interface {
	GetCode() int32
	GetMsg() string
	error
}

type myError struct {
	code Status
	msg  string
}

func (e myError) GetCode() int32 {
	return int32(e.code)
}
func (e myError) GetMsg() string {
	return e.msg
}

func (e myError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

// 数据不存在
func NewDBNoData(msg string) Error {
	if msg == "" {
		msg = StatusDBNotFound.GetMsg()
	}
	return myError{
		code: StatusDBNotFound,
		msg:  msg,
	}
}

//数据重复 420
func NewDBDuplication(key string) Error {

	return myError{
		code: StatusDBDuplication,
		msg:  StatusDBDuplication.GetMsg() + " for " + key,
	}
}

func NewDBInternalErr(err error) Error {

	m := myError{
		code: StatusDBInternalErr,
		msg:  StatusDBInternalErr.GetMsg(),
	}
	if err != nil {
		m.msg=err.Error()
	}
	return m
}
//账户类错误  4001
func NewUnauthorized(msg string) Error {
	if msg == "" {
		msg = StatusUnauthorized.GetMsg()
	}
	return myError{
		code: StatusUnauthorized,
		msg:  msg,
	}
}

//数据权限 4003
func NewForbidden(msg string) Error {

	if msg == "" {
		msg = StatusForbidden.GetMsg()
	}
	return myError{
		code: StatusForbidden,
		msg:  msg,
	}
}

//未知错误 5000
func NewInternalError() Error {
	return &myError{
		code: StatusInternalServerError,
		msg:  StatusInternalServerError.GetMsg(),
	}
}
func InternalErrorMsg(err interface{}) Error {
	var msg string
	if er, ok := err.(error); ok {
		msg = er.Error()
	}
	if er, ok := err.(string); ok {
		msg = er
	}
	return myError{
		code: StatusInternalServerError,
		msg:  msg,
	}
}

// 参数错误 4000
func NewBadRequest(msg string) Error {
	if msg == "" {
		msg = StatusBadRequest.GetMsg()
	}
	return myError{
		code: StatusBadRequest,
		msg:  msg,
	}
}

//database
func New(code Status, err string) Error {
	return myError{code: code, msg: err}
}
