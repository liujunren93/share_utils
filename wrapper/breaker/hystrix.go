package breaker

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/liujunren93/share/wrapper"
	"google.golang.org/grpc"
)

const NAME = "hystrix"

type HystrixRegistry func(menthod string)

func NewClientWrapper(reg HystrixRegistry) wrapper.CallWrapper {
	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			if reg != nil {
				reg(method)
			}

			return hystrix.Do(method, func() error {
				return invoker(ctx, method, req, reply, cc, opts...)
			}, func(err error) error {
				return err
			})

		}, NAME
	}
}
