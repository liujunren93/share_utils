package netHelper

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
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

var (
	Success = HttpResponse{
		Code: 200,
		Msg:  "ok",
		Data: 1,
	}

	//请求参数不匹配
	BindingError = HttpResponse{
		Code: 4000,
		Msg:  "Data verification failed",
	}
	//
	InternalServerError = HttpResponse{
		Code: 5000,
		Msg:  "Internal server error",
	}
	//数据错误
	DataError = HttpResponse{
		Code: 5001,
		Msg:  "Data error",
	}
	OtherError = HttpResponse{
		Code: 0,
		Msg:  "",
	}
)

func (r HttpResponse) GetCode() int32 {
	return int32(r.Code)
}

func (r HttpResponse) GetMsg() string {
	return r.Msg
}

//others[0] status,others[1] data
func Response(o Option, w http.ResponseWriter, msg string, data interface{}) error {

	resData := HttpResponse{
		Code: o.GetCode(),
		Msg:  o.GetMsg(),
		Data: data,
	}
	if msg != "" {
		resData.Msg = msg
	}
	marshal, err := json.Marshal(resData)
	w.WriteHeader(200)
	w.Write(marshal)
	return err
}

//通过反射 设置data
func RpcResponse(a Option, code int32, msg string, data interface{}) error {

	of := reflect.ValueOf(a)
	if of.Kind() != reflect.Ptr && !of.Elem().CanSet() {
		return errors.New("filed")
	}
	elem := of.Elem()
	Data := elem.FieldByName("Data")
	dataOf := reflect.ValueOf(data)

	if dataOf.IsValid() {
		Data.Set(reflect.ValueOf(data))
	}
	elem.FieldByName("Code").SetInt(int64(code))
	elem.FieldByName("Msg").SetString(msg)
	return nil
}
