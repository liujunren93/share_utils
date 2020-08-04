package auth

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	context2 "github.com/shareChina/utils/context"
	"log"
)

func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			shContext := ctx.(*context2.ShContext)
			load, ok := shContext.Header.Load("token")
			log.Fatal(load, ok )
			return fn(ctx, req, rsp)
		}
	}
}