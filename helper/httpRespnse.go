package helper

// 响应
type response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	BindingErr        = 4001 //数据错误
	InternalServerErr = 5000
	DataErr           = 5001
)

func (response) ResSuccess(msg string, data interface{}) *response {
	if msg == "" {
		msg = "ok"
	}
	return newResponse(200, msg, data)
}

func (response) Res5xxErr(code int32, msg string, data interface{}) *response {
	return newResponse(code, msg, data)
}

func newResponse(code int32, msg string, data interface{}) *response {
	return &response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
