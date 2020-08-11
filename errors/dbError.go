package errors

import (
	"github.com/shareChina/utils/helper"
)

type DBError interface {
	Code() helper.Status
	Error() string
}

type dbError struct {
	code helper.Status
	msg  string
}

func (d dbError) Code() helper.Status {
	if d.code == 0 {
		return helper.StatusOK
	}
	return d.code
}

func (d dbError) Error() string {
	return d.msg
}

// 数据不存在
func NoData(msg string) DBError {
	if msg == "" {
		msg = "Data Not Found"
	}
	return &dbError{
		code: 404,
		msg:  msg,
	}
}

//数据重复
func DuplicationData(msg string) DBError {
	if msg == "" {
		msg = "Data Duplication"
	}
	return &dbError{
		code: 503,
		msg:  msg,
	}
}

//未知错误
func DataError(msg string) DBError {
	if msg == "" {
		msg = "Internal Server Error"
	}
	return &dbError{
		code: 500,
		msg:  msg,
	}

}

//database
func NewDBError(code helper.Status, err string) DBError {
	return dbError{code: code, msg: err}
}
