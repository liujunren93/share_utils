package netHelper

import (
	"encoding/json"
	"errors"
	errors2 "github.com/shareChina/utils/errors"
	"github.com/shareChina/utils/helper"
	"net/http"
	"reflect"
)

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//web response
func Response(res helper.StatusI, w http.ResponseWriter, msg string, data interface{}) error {
	resData := HttpResponse{
		Code: res.GetCode(),
		Msg:  res.GetMsg(),
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

//通过反射 设置data rpc response
func RpcResponse(res helper.StatusI, err errors2.Error, data interface{}) error {
	if err == nil {
		err = helper.StatusOK
	}
	of := reflect.ValueOf(res)
	if of.Kind() != reflect.Ptr && !of.Elem().CanSet() {
		return errors.New("filed")
	}
	elem := of.Elem()
	Data := elem.FieldByName("Data")
	dataOf := reflect.ValueOf(data)
	if dataOf.IsValid() {
		Data.Set(reflect.ValueOf(data))
	}
	elem.FieldByName("Code").SetInt(int64(err.GetCode()))
	elem.FieldByName("Msg").SetString(err.GetMsg())
	return nil
}
