package error

import "github.com/shareChina/utils/helper"

//默认数据错误
type dataError struct {
	Code helper.Status //4004 资源不存在,5000 系统异常 5001 sql异常，其余原样输出
	Msg  string
}

func (m dataError) GetMsg() string {
	return m.Msg
}

func (m dataError) GetCode() int32 {
	return int32(m.Code)
}

func (m dataError) Error() string {
	return m.Msg
}
//database
func NewError(code helper.Status, err string) dataError {
	return dataError{Code: code, Msg: err}
}
