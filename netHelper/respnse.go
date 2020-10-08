package netHelper

import (
	"encoding/json"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share_utils/errors"
	"google.golang.org/grpc/status"
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

//Response
func Response(w http.ResponseWriter, sta errors.IStatus, err error, data interface{}) {
	var code int32 = 200
	var msg string = "ok"
	if sta != nil {
		code = sta.GetCode()
		msg = sta.GetMsg()
		if data == nil {
			of := reflect.ValueOf(sta)
			if of.Kind() == reflect.Ptr {
				field := of.Elem().FieldByName("Data")
				if !field.IsZero() {
					data = field.Interface()
				}
			}
		}
	}
	if err != nil {
		msg = err.Error()
	}
	if e, ok := status.FromError(err); err != nil && ok {
		if e.Code() == 400 {
			code = int32(e.Code())
			msg = e.Message()
		} else {
			code = 500
			msg = "Internal Server Error"
		}
	}

	resData := HttpResponse{
		Code: errors.Status(code),
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

//web response
//func Response1(w http.ResponseWriter, sta errors.IStatus, err error, data interface{}) {
//	var code int32
//	var msg string
//	if sta != nil {
//		code = sta.GetCode()
//		msg = sta.GetMsg()
//	}
//	if data == nil {
//		if sta != nil {
//			of := reflect.ValueOf(sta)
//
//			if of.Kind() == reflect.Ptr {
//				field := of.Elem().FieldByName("Data")
//				if !field.IsZero() {
//					data = field.Interface()
//				}
//			}
//
//		}
//	}
//	if _, ok := status.FromError(err); err != nil && ok {
//		code = 500
//		msg = "Internal Server Error"
//	}
//	if err != nil {
//		msg = err.Error()
//	}
//	resData := HttpResponse{
//		Code: errors.Status(code),
//		Msg:  msg,
//		Data: data,
//	}
//	if msg != "" {
//		resData.Msg = msg
//	}
//	marshal, _ := json.Marshal(resData)
//	w.WriteHeader(200)
//	w.Write(marshal)
//
//}

////通过反射 设置data rpc response
func RpcResponse(res errors.IStatus, err errors.Error, data interface{}) error {
	var code int32 = 200
	var msg string = "ok"
	if err != nil {
		if err.GetCode() == 500 {
			return serrors.InternalServerError(err)
		}
		code = err.GetCode()
		msg = err.GetMsg()
	}

	of := reflect.ValueOf(res)
	if of.Kind() != reflect.Ptr && !of.Elem().CanSet() {
		return serrors.InternalServerError(nil)
	}
	elem := of.Elem()
	Data := elem.FieldByName("Data")
	dataOf := reflect.ValueOf(data)
	if dataOf.IsValid() {
		Data.Set(dataOf)
	}
	elem.FieldByName("Code").SetInt(int64(code))
	elem.FieldByName("Msg").SetString(msg)
	return nil
}
