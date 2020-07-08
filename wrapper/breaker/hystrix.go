package breaker

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
)

func NewClientWrapper() client.CallWrapper {
	return func(callFunc client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
				return callFunc(ctx, node, req, rsp, opts)
			}, func(err error) error {
				return err
			})
		}
	}

}
