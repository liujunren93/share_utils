package error

import "github.com/shareChina/utils/helper"

type DataError interface {
	Code()int32
	error
}
type dateError struct {
	Code helper.Status //4004 资源不存在,5000 系统异常 5001 sql异常，其余原样输出
	Msg  string

}

func (m dateError) Error() string {
	return m.Msg
}

//database
func NewError(code helper.Status, err string) *dateError {
	return &dateError{Code: code, Msg: err}
}
