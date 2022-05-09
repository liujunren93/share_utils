package helper

import (
	"fmt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	v3 := GetUuidV3("aaa", true)
	fmt.Println(v3)
}

type aaa struct {
	a string
	A string
}

func TestMapEncodeEscape(t *testing.T) {
	snake := Struct2MapSnakeNoZero(aaa{
		a: "1",
		A: "2",
	})
	fmt.Println(snake)
}

func TestTransSleceType(t *testing.T) {
	var data = []int{1, -2, 3, 4, 5, 6, 7}
	u := TransSliceType[int, uint](data)
	fmt.Println(u)
}
