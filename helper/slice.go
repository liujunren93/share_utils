package helper

import (
	"errors"

	"github.com/liujunren93/share_utils/types"
)

func TransSliceType[OT, NT types.Number](list []OT) []NT {
	var data []NT
	for _, t := range list {

		data = append(data, NT(t))
	}
	return data
}

func InterfaceSlice2NumberSlice[T types.Number](data []interface{}) ([]T, error) {
	var list []T
	for _, da := range data {
		if v, ok := da.(T); ok {
			list = append(list, v)
		} else {
			return nil, errors.New("data type is not []T")
		}
	}

	return list, nil
}
