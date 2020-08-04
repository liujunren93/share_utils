package databases

import "github.com/shareChina/utils/helper"

type DBError interface {
	Code() helper.Status
	Error() string
}

type dbError struct {
	code helper.Status
	msg  string
}

func (d dbError) Code() helper.Status {
	if d.code == 0 {
		return helper.StatusOK
	}
	return d.code
}

func (d dbError) Error() string {
	return d.msg
}

//database
func NewDBError(code helper.Status, err string) DBError {
	return dbError{code: code, msg: err}
}
