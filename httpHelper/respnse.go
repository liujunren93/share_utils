package httpHelper

import (
	"encoding/json"
	"net/http"
)

type Option interface {
	GetCode() int32
	GetMsg() string
	GetData() interface{}
}

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success = HttpResponse{
		Code: 200,
		Msg:  "ok",
		Data: nil,
	}
	BindingError = HttpResponse{
		Code: 4001,
		Msg:  "Data verification failed",
		Data: nil,
	}
	InternalServerError = HttpResponse{
		Code: 5000,
		Msg:  "Internal server error",
		Data: nil,
	}
	DataError = HttpResponse{
		Code: 5001,
		Msg:  "Data error",
		Data: nil,
	}
	OtherError = HttpResponse{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
)

func (r HttpResponse) GetCode() int32 {
	return int32(r.Code)
}

func (r HttpResponse) GetMsg() string {
	return r.Msg
}
func (r HttpResponse) GetData() interface{} {
	return r.Data
}

//others[0] status,others[1] data
func (HttpResponse) Response(o Option, w http.ResponseWriter, others ...interface{}) error {

	marshal, err := json.Marshal(o)
	w.WriteHeader(200)
	w.Write(marshal)
	return err

}
