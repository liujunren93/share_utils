package error

type DataError interface {
	GetCode()int32
	error
}
