package databases

import "github.com/shareChina/utils/helper"

//默认数据错误
type DBStatus struct {
	Code int32
	Msg  string
}

func (m DBStatus) GetCode() int32 {
	return m.Code
}

func (m DBStatus) GetMsg() string {
	return m.Msg
}

func (m DBStatus) Error() string {
	return m.Msg
}

//database
func NewDBError(code helper.Status, err string) DBStatus {
	return DBStatus{Code: int32(code), Msg: err}
}
