package helper

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	BindingErr        = 4001 //数据错误
	InternalServerErr = 5000
	DataErr           = 5001
)

func (HttpResponse) ResSuccess(msg string, data interface{}) *HttpResponse {
	if msg == "" {
		msg = "ok"
	}
	return newResponse(200, msg, data)
}

func (HttpResponse) Res5xxErr(code int32, msg string, data interface{}) *HttpResponse {
	return newResponse(code, msg, data)
}

func newResponse(code int32, msg string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
