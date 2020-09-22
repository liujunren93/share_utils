package metadata

import (
	"context"
	"github.com/liujunren93/share_utils/metadata"
	"google.golang.org/grpc"
)

func ClientUACallWrap(ua *metadata.UserAgent) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		setUA := metadata.SetUA(ctx, ua)
		return invoker(setUA, method, req, reply, cc, opts...)
	}
}
