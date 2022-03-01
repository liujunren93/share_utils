package netHelper

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share/serrors"
	"github.com/liujunren93/share_utils/errors"
	"google.golang.org/grpc/status"
	"reflect"
)

type Responser interface {
	GetCode() int32
	GetMsg() string
}

// 响应
type HttpResponse struct {
	Code errors.Status `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

func ResponseOk(ctx *gin.Context, data interface{})  {
	Response(ctx,errors.StatusOK,nil,data)
}
//Response
//Response
func Response(ctx *gin.Context, res Responser, err error, data interface{}) {

	var code int32 = 200
	var msg = "ok"

	if res != nil {
		code = res.GetCode()
		msg = res.GetMsg()
		if data == nil {
			of := reflect.ValueOf(res)
			if of.Kind() == reflect.Ptr && !of.IsNil() {
				field := of.Elem().FieldByName("Data")
				if field.IsValid() {
					data = field.Interface()
				}
			}
		}
	}

	if err != nil {
		msg = err.Error()
		if e, ok := status.FromError(err); ok {
			if e.Code() >= 5000 {
				code = int32(e.Code())
				msg = e.Message()
			}
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
	ctx.JSON(200, resData)
	ctx.Abort()
}

////通过反射 设置data rpc response
func RpcResponse(res Responser, err errors.Error, data interface{}) error {
	defer func() {
		if errr := recover(); errr != nil {
			errors.InternalErrorMsg(errr)
		}
	}()
	var code int32 = 200
	var msg string = "ok"
	if err != nil {
		if err.GetCode() == 5000 {
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
