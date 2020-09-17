package errors

import (
	"net/http"
)

type IStatus interface {
	GetCode() int32
	GetMsg() (msg string)
}

type Status int32



func (s Status) GetCode() int32 {
	return int32(s)
}

func (s Status) GetMsg() (msg string) {
	if s.GetCode() == 420 {
		return "Data Duplication"
	}
	return http.StatusText(int(s.GetCode()))
}

//database

const (
	StatusOK                  Status = 200 //success
	StatusBadRequest          Status = 400 //数据绑定错误
	StatusUnauthorized        Status = 401 //账户类错误
	StatusForbidden           Status = 403 //权限
	StatusNotFound            Status = 404 //
	StatusDataDuplication     Status = 420 // 数据重复
	StatusInternalServerError Status = 500 //服务器未知错误
)
