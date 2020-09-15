package netHelper

import (
	"encoding/json"
	"github.com/liujunren93/share_utils/errors"
	"net/http"
)

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//web response
func Response(err errors.Error, w http.ResponseWriter, msg string, data interface{}) error {

	var code int32 = 200
	msg = "ok"
	if err != nil {
		code = err.Code()
		msg = err.Error()
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
	return err
}

////通过反射 设置data rpc response
//func RpcResponse(e, err errors2.Error, data interface{}) error {
//	if err == nil {
//		err = utils.StatusOK
//	}
//	of := reflect.ValueOf(res)
//	if of.Kind() != reflect.Ptr && !of.Elem().CanSet() {
//		return errors.New("filed")
//	}
//	elem := of.Elem()
//	Data := elem.FieldByName("Data")
//	dataOf := reflect.ValueOf(data)
//	if dataOf.IsValid() {
//		Data.Set(reflect.ValueOf(data))
//	}
//	elem.FieldByName("Code").SetInt(int64(err.GetCode()))
//	elem.FieldByName("Msg").SetString(err.GetMsg())
//	return nil
//}
