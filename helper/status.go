package helper

type Status int32
func (f Status) Error() string {
	return f.GetMsg()
}

func (f Status) GetCode() int32 {
	return int32(f)
}
const (
	StatusOK                  Status = 200//success
	StatusBadRequest          Status = 4000 //数据绑定错误
	StatusUnauthorized        Status = 4001 //账户类错误
	StatusForbidden           Status = 4003 //权限
	StatusNotFound            Status = 4004 //
	StatusInternalServerError Status = 5000 //服务器未知错误
	StatusDataError           Status = 5001 //database err
)

func (f Status) GetMsg() (msg string) {

	switch f {
		case StatusOK:
			msg = "ok"
		case StatusBadRequest:
			msg = "Request Data Error"
		case StatusUnauthorized:
			msg = "Status Unauthorized"
		case StatusForbidden:
			msg = "Status Forbidden"
		case StatusNotFound:
			msg = "Status Not Found"
		case StatusInternalServerError:
			msg = "Status Internal Server Error"
		case StatusDataError:
			msg = "status Data Error"
		default:

	}
	return
}
