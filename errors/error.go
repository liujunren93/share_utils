package errors

import (
	"fmt"
	"strings"
)

type Error interface {
	GetCode() int32
	GetMsg() string
	error
}

type ShError struct {
	code Status
	msg  string
}

func (e ShError) GetCode() int32 {
	return int32(e.code)
}
func (e ShError) GetMsg() string {
	return e.msg
}

func (e ShError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

// 数据不存在
func NewDBNoData(msg string) Error {
	if msg == "" {
		msg = StatusDBNotFound.GetMsg()
	}
	return ShError{
		code: StatusDBNotFound,
		msg:  msg,
	}
}

//数据重复 420
func NewDBDuplication(key string) Error {

	return ShError{
		code: StatusDBDuplication,
		msg:  StatusDBDuplication.GetMsg() + " for " + key,
	}
}

func NewDBInternal(err error) Error {

	m := ShError{
		code: StatusDBInternalErr,
		msg:  StatusDBInternalErr.GetMsg(),
	}
	if err != nil {
		m.msg = err.Error()
	}
	return m
}

//账户类错误  4001
func NewUnauthorized(msg string) Error {
	if msg == "" {
		msg = StatusUnauthorized.GetMsg()
	}
	return ShError{
		code: StatusUnauthorized,
		msg:  msg,
	}
}

//数据权限 4003
func NewForbidden(msg string) Error {

	if msg == "" {
		msg = StatusForbidden.GetMsg()
	}
	return ShError{
		code: StatusForbidden,
		msg:  msg,
	}
}

//未知错误 5000
func NewInternalError() Error {
	return &ShError{
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
	return ShError{
		code: StatusInternalServerError,
		msg:  msg,
	}
}

// 参数错误 4000
func NewBadRequest(msg string) Error {
	if msg == "" {
		msg = StatusBadRequest.GetMsg()
	}
	return ShError{
		code: StatusBadRequest,
		msg:  msg,
	}
}

//database
func New(code Status, err string) Error {
	return ShError{code: code, msg: err}
}

type ShErrors struct {
	code Status
	Errs []Error
}

func (e ShErrors) GetCode() int32 {
	return int32(e.code)
}
func (e ShErrors) GetMsg() string {
	return e.Error()
}

func (e ShErrors) Error() string {
	var b = strings.Builder{}
	for _, e := range e.Errs {
		b.WriteString(e.Error())
	}
	return b.String()
}

func (e *ShErrors) AddErr(code Status, err string) {
	e.Errs = append(e.Errs, New(code, err))
}

func NewS(code Status) ShErrors {
	return ShErrors{}
}
