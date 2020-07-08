package breaker

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
)

func NewClientWrapper(maxConcurrent, Timeout int) client.CallWrapper {

	if maxConcurrent != 0 {
		hystrix.DefaultMaxConcurrent = maxConcurrent
	}
	if Timeout != 0 {
		hystrix.DefaultTimeout = Timeout
	}
	hystrix.DefaultVolumeThreshold = 2
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
