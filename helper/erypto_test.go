package helper

import (
	"fmt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	v3 := GetUuidV3("aaa", true)
	fmt.Println(v3)
}

func TestMapEncodeEscape(t *testing.T) {
	escape := MapEncodeEscape(map[string]interface{}{"a": 1, "b": 1.2, "c": "hhh", "d": rune(1)})
	fmt.Println(escape)

}
