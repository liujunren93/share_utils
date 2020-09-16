package errors

type Error interface {
	GetCode() int32
	GetMsg() string
	error
}

type myError struct {
	Code Status
	Msg  string
}

func (e myError) GetCode() int32 {
	return int32(e.Code)
}
func (e myError) GetMsg() string {
	return e.Msg
}

func (e myError) Error() string {
	return e.Msg
}

// 数据不存在
func NoData(msg string) Error {
	if msg == "" {
		msg = StatusNotFound.GetMsg()
	}
	return &myError{
		Code: StatusNotFound,
		Msg:  msg,
	}
}

//数据重复 420
func DuplicationData(msg string) Error {
	if msg == "" {
		msg = StatusDataDuplication.GetMsg()
	}
	return &myError{
		Code: StatusDataDuplication,
		Msg:  msg,
	}
}

//账户类错误  401
func Unauthorized(msg string) Error {
	if msg == "" {
		msg = StatusUnauthorized.GetMsg()
	}
	return &myError{
		Code: StatusUnauthorized,
		Msg:  msg,
	}
}

//数据权限 403
func Forbidden(msg string) Error {

	if msg == "" {
		msg = StatusForbidden.GetMsg()
	}
	return &myError{
		Code: StatusForbidden,
		Msg:  msg,
	}
}

//未知错误 500
func DataError(msg string) Error {
	if msg == "" {
		msg = StatusInternalServerError.GetMsg()
	}
	return &myError{
		Code: StatusInternalServerError,
		Msg:  msg,
	}

}

// 参数错误 400
func BadRequest(msg string) Error {
	if msg == "" {
		msg = StatusBadRequest.GetMsg()
	}
	return &myError{
		Code: StatusBadRequest,
		Msg:  msg,
	}
}

//database
func New(code Status, err string) Error {
	return myError{Code: code, Msg: err}
}
