package helper

type Status int32

func (f Status) GetCode() int32 {
	return int32(f)
}

const (
	StatusOK                  Status = 200
	StatusBadRequest          Status = 4000
	StatusUnauthorized        Status = 4001
	StatusForbidden           Status = 4003
	StatusNotFound            Status = 4004
	StatusInternalServerError Status = 5000
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
