package auth

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
)

func NewClientAuthWrapper(token string) client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {

			ctx = metadata.Set(ctx, "Authorization", token)

			return cf(ctx, node, req, rsp, opts)
		}
	}
}
