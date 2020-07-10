package netHelper

import (
	"fmt"
	"reflect"
	"testing"
)

type test interface {
	getname()
}

type User struct {
	Name string
	ID   int
	Age  int
}

func (User) getname() {

}
func TestRpcResponse(t *testing.T) {
	a := 1


}

