package netHelper

import (
	"encoding/json"
	"github.com/liujunren93/share_utils/errors"
	"net/http"
)

// 响应
type HttpResponse struct {
	Code errors.Status       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//web response
func Response(w http.ResponseWriter,code errors.Status, msg string, data interface{})  {
	msg = "ok"
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
