package databases

import "github.com/shareChina/utils/helper"

type DBError interface {
	Code() helper.Status
	Msg() string
}

type dbError struct {
	code helper.Status
	msg  string
}

func (d dbError) Code() helper.Status {
	if d.code == 0 {
		return 200
	}
	return d.code
}

func (d dbError) Msg() string {
	return d.msg
}

//database
func NewDBError(code helper.Status, err string) DBError {
	return dbError{code: code, msg: err}
}
