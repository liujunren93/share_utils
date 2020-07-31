package netHelper

type fixedCode int32

func (f fixedCode) GetCode() int32 {
	return int32(f)
}

const (
	StatusOK                  fixedCode = 200
	StatusBadRequest          fixedCode = 400
	StatusUnauthorized        fixedCode = 4001
	StatusForbidden           fixedCode = 4003
	StatusNotFound            fixedCode = 4004
	StatusInternalServerError fixedCode = 5000
	statusDataError           fixedCode = 5001
)

func (f fixedCode) GetMsg() (msg string) {

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
	case statusDataError:
		msg = "status Data Error"
	}
	return
}
