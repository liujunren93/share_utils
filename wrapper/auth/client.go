package auth

import (
	"context"

	"github.com/liujunren93/share/wrapper"
	"google.golang.org/grpc"
)

const CLIENT_NAME = "auth_client"

func NewClientWrapper(key string, f func(context.Context) interface{}) wrapper.CallWrapper {
	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			// ctx, err := metadata.SetVal(ctx, key,)
			// if err != nil {
			// 	return err
			// }
			//TODO set auth
			return invoker(ctx, method, req, reply, cc, opts...)
		}, CLIENT_NAME
	}
}
