package metadata

import (
	"context"

	"github.com/liujunren93/share/wrapper"
	"github.com/liujunren93/share_utils/common/metadata"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const NAME = "metadata"

func NewClientWrapper(key string, f func(context.Context) ([]byte, error)) wrapper.CallWrapper {

	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

			data, err := f(ctx)
			if err != nil {
				return err
			}
			ctx, err = metadata.SetVal(ctx, key, string(data))
			if err != nil {
				return err
			}
			return invoker(ctx, method, req, reply, cc, opts...)
		}, NAME + key
	}

}

func NewClientWrapperMessage(key string, f func(context.Context) (proto.Message, error)) wrapper.CallWrapper {

	return func() (grpc.UnaryClientInterceptor, string) {
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			val, err := f(ctx)
			if err != nil {
				return err
			}
			ctx, err = metadata.SetMessage(ctx, key, val)
			if err != nil {
				return err
			}
			return invoker(ctx, method, req, reply, cc, opts...)
		}, NAME + key
	}

}
