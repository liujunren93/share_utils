package breaker

import (
	"context"
	"fmt"
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

func NewClientWrapper(maxConcurrent, Timeout int) client.Wrapper {
	fmt.Println(111111)
	if maxConcurrent == 0 {
		hystrix.DefaultMaxConcurrent = 10
	} else {
		hystrix.DefaultMaxConcurrent = maxConcurrent
	}
	if Timeout == 0 {
		hystrix.DefaultTimeout = 1000
	}else{
		hystrix.DefaultTimeout = Timeout
	}


	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
