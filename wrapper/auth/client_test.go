package auth

import (
	"fmt"
	"github.com/shareChina/utils/context"
	"testing"
)

func TestNewClientAuthWrapper(t *testing.T) {
	newContext := context.NewContext()
	newContext.Header.Store("aaa","asad")
	load, ok := newContext.Header.Load("aaa")
	fmt.Println(load,ok)
}