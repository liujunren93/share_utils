package errors

import "fmt"

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
	return fmt.Sprintf("code:%d,msg:%s",e.code,e.msg)
}

// 数据不存在
func NoData(msg string) Error {
	if msg == "" {
		msg = StatusNotFound.GetMsg()
	}
	return &myError{
		code: StatusNotFound,
		msg:  msg,
	}
}

//数据重复 420
func DuplicationData(msg string) Error {
	if msg == "" {
		msg = StatusDataDuplication.GetMsg()
	}
	return &myError{
		code: StatusDataDuplication,
		msg:  msg,
	}
}

//账户类错误  401
func Unauthorized(msg string) Error {
	if msg == "" {
		msg = StatusUnauthorized.GetMsg()
	}
	return &myError{
		code: StatusUnauthorized,
		msg:  msg,
	}
}

//数据权限 403
func Forbidden(msg string) Error {

	if msg == "" {
		msg = StatusForbidden.GetMsg()
	}
	return &myError{
		code: StatusForbidden,
		msg:  msg,
	}
}

//未知错误 500
func DataError(msg string) Error {
	if msg == "" {
		msg = StatusInternalServerError.GetMsg()
	}
	return &myError{
		code: StatusInternalServerError,
		msg:  msg,
	}

}

// 参数错误 400
func BadRequest(msg string) Error {
	if msg == "" {
		msg = StatusBadRequest.GetMsg()
	}
	return &myError{
		code: StatusBadRequest,
		msg:  msg,
	}
}

func Timeout(msg string)Error  {
	if msg == "" {
		msg = "time out"
	}
	return &myError{
		code: 408,
		msg:  msg,
	}
}



//database
func New(code Status, err string) Error {
	return myError{code: code, msg: err}
}
