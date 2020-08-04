package auth

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	context2 "github.com/shareChina/utils/context"
)

func NewClientAuthWrapper(token string) client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			todo := context2.Todo
			todo.Header.Store("token", token)
			return cf(todo, node, req, rsp, opts)
		}
	}
}
