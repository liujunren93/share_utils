package httpHandler

import (
	"context"
	"fmt"
	"testing"
)


func Test_rpcRequest(t *testing.T) {
	request, err := NewRPCRequest("aaa", "dada", "sdas", "", "")
	rpc, err := request.RPC(context.Background())

	fmt.Println(rpc,err)
}
type str string
func TestBb(t *testing.T) {

}