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

func TestAesEncrypt(t *testing.T) {
	str, err := AesEncrypt("gGJk6duYH0ucfxGA1Bi9Ag1aMyFDiAFt9glaDnHFzKpEsMfddLviMQyau6y2ji0g", "liujunrenliujun1")
	fmt.Println(str, err)
	str, err = AesDecrypt(str, "liujunrenliujun1")
	fmt.Println(str, err)
}

// func TestAesDecrypt(t *testing.T) {
// 	str := AesEncrypt1("liujunrenliujunr", "liujunrenliujunr")
// 	fmt.Println(str)
// 	str = AesDecrypt1(str, "liujunrenliujunr")
// 	fmt.Println(str)
// }
