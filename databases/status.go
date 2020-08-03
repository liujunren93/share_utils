package databases

import "github.com/shareChina/utils/helper"

type ModelError struct {
	Code helper.Status //4004 资源不存在,5000 系统异常 5001 sql异常，其余原样输出
	Msg  string
}



func (m ModelError) Error() string {
	return m.Msg
}

func (Base) NewError(code helper.Status, err string) *ModelError {
	return &ModelError{Code: code, Msg: err}
}

