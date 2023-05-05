package recover

import (
	"context"
	"fmt"
	"runtime"

	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/log"
	"google.golang.org/grpc"
)

func Recover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if re := recover(); re != nil {
				resp = nil
				err = errors.NewInternalError(re)
				buf := make([]byte, 1<<16)
				len := runtime.Stack(buf, true)
				fmt.Println("recover", string(buf[:len]))
				log.Logger.Error("wrapper recover", string(buf[:len]))
			}
		}()
		return handler(ctx, req)
	}
}
