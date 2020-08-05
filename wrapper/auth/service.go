package auth

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"log"
)

func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {

			get, b := metadata.Get(ctx, "Authorization")
			log.Fatal(get, b)
			return fn(ctx, req, rsp)
		}
	}
}
