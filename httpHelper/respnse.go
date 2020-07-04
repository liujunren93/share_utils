package httpHelper

import (
	"encoding/json"
	"errors"
	"github.com/shareChina/utils/log"
	"net/http"
)

type Option interface {
	GetStatus() (int32, string)
}

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// success
type (
	Success             string
	BindingError        string
	InternalServerError string
	DataError           string
	OtherError          string
)

// 数据校验不通过
func (e Success) getStatus() (int32, string) {
	if e == "" {
		e = "ok"
	}
	return 200, string(e)
}

func (e BindingError) GetStatus() (int32, string) {
	if e == "" {
		e = "Data verification failed"
	}
	return 4001, string(e)
}

func (e InternalServerError) GetStatus() (int32, string) {
	if e == "" {
		e = "Internal server error"
	}
	return 5000, string(e)
}

func (e DataError) GetStatus() (int32, string) {
	if e == "" {
		e = "Data error"
	}
	return 5001, string(e)
}

func (e OtherError) GetStatus() (int32, string) {
	if e != "" {
		e = "Unknown mistake"
	}
	return 0, string(e)
}

//others[0] status,others[1] data
func (HttpResponse) Response(o Option, w http.ResponseWriter, others ...interface{}) error {

	status, msg := o.GetStatus()
	if status == 0 && len(others) == 0 && others[0] != nil {
		log.Logger.Fatal("you must give an  error code ")
		return errors.New("you must give an  error code ")
	}
	var data interface{}
	if len(others) > 0 && others[0] != nil {
		status = int32(others[0].(int))
	}
	if len(others) > 1 && others[1] != nil {
		data = others[1]
	}

	response := newResponse(status, msg, data)
	marshal, err := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(marshal)
	return err

}

func newResponse(code int32, msg string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
