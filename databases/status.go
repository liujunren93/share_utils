package databases

type ModelError struct {
	Code int32 //4004 资源不存在,5000 系统异常 5001 sql异常，其余原样输出
	Msg  string
}

var (
	DataErr             = &ModelError{Code: 5001, Msg: "data error"}
	InternalServerError = &ModelError{Code: 5000, Msg: "Internal Server Error"}
	NotFound            = &ModelError{Code: 4004, Msg: "数据不存在"}
)

func (m ModelError) Error() string {
	return m.Msg
}

func (Base) NewError(code int32, err string) *ModelError {
	return &ModelError{Code: code, Msg: err}
}

