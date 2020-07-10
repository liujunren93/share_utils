package netHelper

import (
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
	RpcResponse(&Success,200,"sss",nil)


}

