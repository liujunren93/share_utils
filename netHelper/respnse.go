package netHelper

import (
	"encoding/json"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share_utils/errors"
	"net/http"
	"reflect"
)

type res interface {
	GetMsg() int32
	GetCode() string
}

// 响应
type HttpResponse struct {
	Code errors.IStatus `json:"code"`
	Msg  string         `json:"msg"`
	Data interface{}    `json:"data"`
}

//func Response5XX(w http.ResponseWriter, err error, msg string, data interface{}) {
//	if msg == "" {
//		msg = err.Error()
//	}
//	ResponseOK(w, errors.StatusInternalServerError, msg, data)
//}

//web response
func Response(w http.ResponseWriter, code errors.IStatus, msg string, data interface{}) {
	if msg == "" {
		msg = code.GetMsg()
	}
	resData := HttpResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	if msg != "" {
		resData.Msg = msg
	}
	marshal, _ := json.Marshal(resData)
	w.WriteHeader(200)
	w.Write(marshal)

}

////通过反射 设置data rpc response
func RpcResponse(res interface{}, err errors.Error, data interface{}) (interface{}, error) {
	if err.GetCode() == 500 {
		return res, serrors.InternalServerError(err)
	}
	of := reflect.ValueOf(res)
	if of.Kind() != reflect.Ptr && !of.Elem().CanSet() {
		return res, serrors.InternalServerError(err)
	}
	elem := of.Elem()
	Data := elem.FieldByName("Data")
	dataOf := reflect.ValueOf(data)
	if dataOf.IsValid() {
		Data.Set(reflect.ValueOf(data))
	}
	elem.FieldByName("Code").SetInt(int64(err.GetCode()))
	elem.FieldByName("Msg").SetString(err.GetMsg())
	return res, nil
}
