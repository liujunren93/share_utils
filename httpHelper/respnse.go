package httpHelper

import (
	"encoding/json"
	"net/http"
)

type Option interface {
	GetCode() int32
	GetMsg() string
}

// 响应
type httpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success = httpResponse{
		Code: 200,
		Msg:  "ok",
	}
	BindingError = httpResponse{
		Code: 4001,
		Msg:  "Data verification failed",
	}
	InternalServerError = httpResponse{
		Code: 5000,
		Msg:  "Internal server error",
	}
	DataError = httpResponse{
		Code: 5001,
		Msg:  "Data error",
	}
	OtherError = httpResponse{
		Code: 0,
		Msg:  "",
	}
)

func (r httpResponse) GetCode() int32 {
	return int32(r.Code)
}

func (r httpResponse) GetMsg() string {
	return r.Msg
}

//others[0] status,others[1] data
func Response(o Option, w http.ResponseWriter, data interface{}) error {
	resData := httpResponse{
		Code: o.GetCode(),
		Msg:  o.GetMsg(),
		Data: data,
	}
	marshal, err := json.Marshal(resData)
	w.WriteHeader(200)
	w.Write(marshal)
	return err

}
