package breaker

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, nil)
}

func NewClientWrapper() client.Wrapper {
	hystrix.DefaultMaxConcurrent = 3
	hystrix.DefaultTimeout = 2
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
