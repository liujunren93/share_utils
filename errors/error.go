package errors

import (
	"github.com/liujunren93/share_utils"
)

type Error interface {
	Code()int32
	error
}

type myError struct {
	code utils.Status
	msg  string
}

func (e myError) Code()int32 {
	return int32(e.code)
}


func (e myError) Error() string {
	return e.msg
}

// 数据不存在
func NoData(msg string) Error {
	if msg == "" {
		msg = utils.StatusNotFound.GetMsg()
	}
	return &myError{
		code: utils.StatusNotFound,
		msg:  msg,
	}
}

//数据重复 420
func DuplicationData(msg string) Error {
	if msg == "" {
		msg = utils.StatusDataDuplication.GetMsg()
	}
	return &myError{
		code: utils.StatusDataDuplication,
		msg:  msg,
	}
}

//账户类错误  401
func Unauthorized(msg string) Error {
	if msg == "" {
		msg = utils.StatusUnauthorized.GetMsg()
	}
	return &myError{
		code: utils.StatusUnauthorized,
		msg:  msg,
	}
}

//数据权限 403
func Forbidden(msg string) Error {

	if msg == "" {
		msg = utils.StatusForbidden.GetMsg()
	}
	return &myError{
		code: utils.StatusForbidden,
		msg:  msg,
	}
}

//未知错误 500
func DataError(msg string) Error {
	if msg == "" {
		msg = utils.StatusInternalServerError.GetMsg()
	}
	return &myError{
		code: utils.StatusInternalServerError,
		msg:  msg,
	}

}

// 参数错误 400
func BadRequest(msg string) Error {
	if msg == "" {
		msg = utils.StatusBadRequest.GetMsg()
	}
	return &myError{
		code: utils.StatusBadRequest,
		msg:  msg,
	}
}

//database
func New(code utils.Status, err string) Error {
	return myError{code: code, msg: err}
}
