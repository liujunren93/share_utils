package helper

type StatusI interface {
	GetCode() int32
	GetMsg() string
}

type Status int32

func (s Status) Error() string {
	return s.GetMsg()
}

func (s Status) GetCode() int32 {
	return int32(s)
}

func (s Status) GetMsg() (msg string) {
	switch s {
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
	case StatusDataDuplication:
		msg = "Data Duplication"
	case StatusInternalServerError:
		msg = "Status Internal Server Error"
	case StatusDataError:
		msg = "Status Data Error"
	default:
		msg = "Unknown Mistake"
	}
	return
}

//database

const (
	StatusOK                  Status = 200  //success
	StatusBadRequest          Status = 400 //数据绑定错误
	StatusUnauthorized        Status = 401 //账户类错误
	StatusForbidden           Status = 403 //权限
	StatusNotFound            Status = 404 //
	StatusDataDuplication     Status = 405 // 数据重复
	StatusInternalServerError Status = 500 //服务器未知错误
	StatusDataError           Status = 501 //database err
)
