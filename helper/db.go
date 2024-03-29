package helper

import (
	"fmt"
)

func PageSize(size int32) int {
	if size <= 0 || size > 50 {
		return 25
	}
	return int(size)
}

func Offset(size, page int32) int {
	size = int32(PageSize(size))
	return int(size * page)
}

func LikeLeft(field string) string {
	return fmt.Sprintf("%%%v", field)
}

func LikeRight(field string) string {
	return fmt.Sprintf("%v%%", field)
}

func Like(field string) string {
	return fmt.Sprintf("%%%v%%", field)
}

func Order(fieldorOrders ...string) map[string]string {
	var order = make(map[string]string)
	for i := 0; i < len(fieldorOrders); i += 2 {
		order[fieldorOrders[i]] = fieldorOrders[i+1]
	}

	return order
}
