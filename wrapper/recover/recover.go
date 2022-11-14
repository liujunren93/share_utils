package recover

import (
	"context"
	"fmt"

	"github.com/liujunren93/share_utils/errors"
	"google.golang.org/grpc"
)

func Recover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if re := recover(); re != nil {
				resp = nil
				err = errors.NewInternalError(re)
				fmt.Println(resp, err)
			}
		}()
		return handler(ctx, req)
	}
}
