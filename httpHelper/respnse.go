package httpHelper

type option interface {
	getStatus() (int32, string)
}

// 响应
type HttpResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 数据校验不通过
type BindingError string

func (e BindingError) getStatus() (int32, string) {
	if e != "" {
		e = "数据校验不通过"
	}
	return 4001, string(e)
}

type InternalServerError string

func (e InternalServerError) getStatus() (int32, string) {
	if e != "" {
		e = "Internal Server Error"
	}
	return 5000, string(e)
}

type DataError string

func (e DataError) getStatus() (int32, string) {
	if e != "" {
		e = "服务器数据错误"
	}
	return 5000, string(e)
}

type OtherError string

func (e OtherError) getStatus() (int32, string) {
	if e != "" {
		e = "未知错误"
	}
	return 5000, string(e)
}

func (HttpResponse) ResSuccess(o option, data interface{}) *HttpResponse {
	status, msg := o.getStatus()
	return newResponse(status, msg, data)
}

func (HttpResponse) ResError(o option, others ...interface{}) *HttpResponse {
	status, msg := o.getStatus()
	if others[0] != nil {
		status = others[0].(int32)
	}
	return newResponse(status, msg, others[1])
}

func newResponse(code int32, msg string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
