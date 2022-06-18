package types

type Number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
