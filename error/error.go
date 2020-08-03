package error

type DataError interface {
	Code()int32
	error
}
