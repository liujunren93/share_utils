package errors

import (
	"github.com/shareChina/utils"
)

type Error interface {
	utils.StatusI
}

type error struct {
	code utils.Status
	msg  string
}

func (d error) GetCode() int32 {
	if d.code == 0 {
		return 200
	}
	return int32(d.code)
}

func (d error) GetMsg() string {
	return d.msg
}

// 数据不存在
func NoData(msg string) Error {
	if msg == "" {
		msg = utils.StatusNotFound.GetMsg()
	}
	return &error{
		code: utils.StatusNotFound,
		msg:  msg,
	}
}

//数据重复
func DuplicationData(msg string) Error {
	if msg == "" {
		msg = utils.StatusDataDuplication.GetMsg()
	}
	return &error{
		code: utils.StatusDataDuplication,
		msg:  msg,
	}
}

//账户类错误
func Unauthorized(msg string) Error {
	if msg == "" {
		msg = utils.StatusUnauthorized.GetMsg()
	}
	return &error{
		code: utils.StatusUnauthorized,
		msg:  msg,
	}
}

//数据权限
func Forbidden(msg string) Error {
	if msg == "" {
		msg = utils.StatusForbidden.GetMsg()
	}
	return &error{
		code: utils.StatusForbidden,
		msg:  msg,
	}
}



//未知错误
func DataError(msg string) Error {
	if msg == "" {
		msg = utils.StatusInternalServerError.GetMsg()
	}
	return &error{
		code: utils.StatusInternalServerError,
		msg:  msg,
	}

}

//database
func New(code utils.Status, err string) Error {
	return error{code: code, msg: err}
}

