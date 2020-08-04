package databases

import "github.com/shareChina/utils/helper"

//默认数据错误
type dataStatus struct {
	Code helper.Status
	Msg  string
}

func (m dataStatus) GetCode() int32 {
	return int32(m.Code)
}

func (m dataStatus) GetMsg() string {
	return m.Msg
}

func (m dataStatus) Error() string {
	return m.Msg
}
//database
func NewError(code helper.Status, err string) dataStatus {
	return dataStatus{Code: code, Msg: err}
}
