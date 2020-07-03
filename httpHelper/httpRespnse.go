package httpHelper

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	BindingErr        int32 = 4001 //数据错误
	InternalServerErr int32 = 5000
	DataErr           int32 = 5001
)

func (HttpResponse) ResSuccess(msg string, data interface{}) *HttpResponse {
	if msg == "" {
		msg = "ok"
	}
	return newResponse(200, msg, data)
}

func (HttpResponse) ResError(code int32, msg string, data interface{}) *HttpResponse {
	return newResponse(code, msg, data)
}

func newResponse(code int32, msg string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
