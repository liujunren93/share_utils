package errors

import (
	"github.com/shareChina/utils/netHelper"
)

type Error interface {
	GetCode() int32
	GetMsg() string
}

type error struct {
	code netHelper.Status
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
		msg = netHelper.StatusNotFound.GetMsg()
	}
	return &error{
		code: netHelper.StatusNotFound,
		msg:  msg,
	}
}

//数据重复
func DuplicationData(msg string) Error {
	if msg == "" {
		msg = netHelper.StatusDataDuplication.GetMsg()
	}
	return &error{
		code: netHelper.StatusDataDuplication,
		msg:  msg,
	}
}

//未知错误
func DataError(msg string) Error {
	if msg == "" {
		msg = netHelper.StatusInternalServerError.GetMsg()
	}
	return &error{
		code: netHelper.StatusInternalServerError,
		msg:  msg,
	}

}

//database
func New(code netHelper.Status, err string) Error {
	return error{code: code, msg: err}
}

