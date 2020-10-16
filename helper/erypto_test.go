package helper

import (
	"fmt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	v3 := GetUuidV3("aaa",true)
	fmt.Println(v3)
}

