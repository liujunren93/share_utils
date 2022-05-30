package metadata

import (
	"context"

	"github.com/liujunren93/share/wrapper"
	"github.com/liujunren93/share_utils/common/metadata"
	"google.golang.org/grpc"
)

const NAME = "metadata"

func NewClientWrapper(key string, f func(context.Context) string) wrapper.CallWrapper {

	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			ctx, err := metadata.SetVal(ctx, key, f(ctx))
			if err != nil {
				return err
			}
			panic(111)
			return invoker(ctx, method, req, reply, cc, opts...)
		}, NAME + key
	}

}
