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
	return NewPublic(StatusDBNotFound, msg)
}

//数据重复 420
func NewDBDuplication(key string) Error {
	return NewPublic(StatusDBDuplication, StatusDBDuplication.GetMsg()+" for "+key)
}

func NewDBInternal(err error) Error {

	return New(StatusDBInternalErr, err)
}

//账户类错误  4001
func NewUnauthorized(msg string) Error {
	return NewPublic(StatusUnauthorized, msg)
}

//数据权限 4003
func NewForbidden(msg interface{}) Error {

	return NewPublic(StatusForbidden, msg)
}

//未知错误 5000
func NewInternalError(err interface{}) Error {

	return New(StatusInternalServerError, err)
}

// 参数错误 4000
func NewBadRequest(err interface{}) Error {

	return NewPublic(StatusBadRequest, err)
}

//database
func New(code Status, err interface{}) Error {
	m := getMsg(code, err)
	return ShError{code: code * 10, msg: m}
}

func NewPublic(code Status, err interface{}) Error {
	m := getMsg(code, err)
	return ShError{code: code*10 + 1, msg: m}
}
func getMsg(s Status, err interface{}) string {
	msg := s.GetMsg()
	if err != nil {
		switch m := err.(type) {
		case string:
			if m != "" {
				msg = m
			}

		case error:
			msg = m.Error()
		default:
			msg = fmt.Sprintf("%v", m)
		}
	}

	return msg
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
