package netHelper

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

type Return interface {
	GetCode() int32
	GetMsg() string
}



// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}



func (r HttpResponse) GetCode() int32 {
	return int32(r.Code)
}

func (r HttpResponse) GetMsg() string {
	return r.Msg
}

//others[0] status,others[1] data
func Response(r Return, w http.ResponseWriter, msg string, data interface{}) error {

	resData := HttpResponse{
		Code: r.GetCode(),
		Msg:  r.GetMsg(),
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
func RpcResponse(r Return, code int32, msg string, data interface{}) error {

	of := reflect.ValueOf(r)
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
