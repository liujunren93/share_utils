package httpHelper

import (
	"encoding/json"
	"errors"
	"github.com/shareChina/utils/log"
	"net/http"
)

type Option interface {
	GetCode() int32
	GetMsg() string
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

// success
func (e Success) GetCode() int32 {
	return 200
}

func (e Success) GetMsg() string {
	if e == "" {
		e = "ok"
	}
	return string(e)
}

// binding err
func (e BindingError) GetCode() int32 {
	return 4001
}

func (e BindingError) GetMsg() string {
	if e == "" {
		e = "Data verification failed"
	}
	return string(e)
}

// Internal server error
func (e InternalServerError) GetCode() int32 {
	return 5000
}

func (e InternalServerError) GetMsg() string {
	if e == "" {
		e = "Internal server error"
	}
	return string(e)
}

// Data error
func (e DataError) GetCode() int32 {
	return 5001
}

func (e DataError) GetMsg() string {
	if e == "" {
		e = "Data error"
	}
	return string(e)
}

// Unknown mistake
func (e OtherError) GetCode() int32 {
	return 0
}

func (e OtherError) GetMsg() string {
	if e == "" {
		e = "Unknown mistake"
	}
	return string(e)
}

//others[0] status,others[1] data
func (HttpResponse) Response(o Option, w http.ResponseWriter, others ...interface{}) error {

	status, msg := o.GetCode(), o.GetMsg()
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
