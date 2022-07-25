package metadata

import (
	"context"
	"fmt"

	"github.com/liujunren93/share/wrapper"
	"github.com/liujunren93/share_utils/common/metadata"
	"google.golang.org/grpc"
)

const NAME = "metadata"

func NewClientWrapper(key string, f func(context.Context) string) wrapper.CallWrapper {

	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			fmt.Println(f(ctx))
			ctx, err := metadata.SetVal(ctx, key, f(ctx))
			if err != nil {
				return err
			}

			ctx, _ = metadata.SetVal(ctx, "aaaa", "111")
			fmt.Println(metadata.GetAll(ctx))
			return invoker(ctx, method, req, reply, cc, opts...)
		}, NAME + key
	}

}
