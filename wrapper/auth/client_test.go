package auth

import (
	"context"
	"fmt"
	c2 "github.com/shareChina/utils/context"
	"testing"
)

func TestNewClientAuthWrapper(t *testing.T) {
	var a context.Context
	newContext := c2.NewContext()
	newContext.Header.Store("aaa","asad")
	a=newContext
	//load, ok := newContext.Header.Load("aaa")

	shContext := a.(*c2.ShContext)
	fmt.Println(shContext.Header.Load("aaa"))
}