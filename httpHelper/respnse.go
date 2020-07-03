package httpHelper

import (
	"encoding/json"
	"errors"
	"github.com/shareChina/utils/log"
	"net/http"
)

type Option interface {
	getStatus() (int32, string)
}

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// success
type Success string

// 数据校验不通过
func (e Success) getStatus() (int32, string) {
	if e != "" {
		e = "ok"
	}
	return 200, string(e)
}

type BindingError string

func (e BindingError) getStatus() (int32, string) {
	if e != "" {
		e = "Data verification failed"
	}
	return 4001, string(e)
}

type InternalServerError string

func (e InternalServerError) getStatus() (int32, string) {
	if e != "" {
		e = "Internal server error"
	}
	return 5000, string(e)
}

type DataError string

func (e DataError) getStatus() (int32, string) {
	if e != "" {
		e = "Data error"
	}
	return 5001, string(e)
}

type OtherError string

func (e OtherError) getStatus() (int32, string) {
	if e != "" {
		e = "Unknown mistake"
	}
	return 0, string(e)
}

func (HttpResponse) Response(o Option, w http.ResponseWriter, others ...interface{}) error {
	status, msg := o.getStatus()
	if status == 0 && others[0] != nil {
		log.Logger.Fatal("you must give an  error code ")
		return errors.New("you must give an  error code ")
	}
	if others[0] != nil {
		status = others[0].(int32)
	}
	response := newResponse(status, msg, others[1])
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
