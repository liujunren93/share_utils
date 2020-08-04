package databases

import "github.com/shareChina/utils/helper"



type DBError struct {
	Code helper.Status
	Msg  string
}

//database
func NewDBError(code helper.Status, err string) *DBError {
	return &DBError{Code: code, Msg: err}
}
