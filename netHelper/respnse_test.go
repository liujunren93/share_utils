package netHelper

import (
	"fmt"
	"reflect"
	"testing"
)

type args struct {
	Code int32
	Msg  string
	Data interface{}
}

func (a *args) GetCode() int32 {

	return 0
}

func (a *args) GetMsg() string {
	panic("implement me")
}


func aaa(a interface{}) () {
	of := reflect.ValueOf(a)
	fmt.Println(of.IsZero())

}
func TestRpcResponse(t *testing.T) {
	return

}
