package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

//获取
func getValue(ctx context.Context, key string) ([]string, bool) {
	incomingContext, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ok
	}
	strings, ok := incomingContext[key]
	return strings, ok
}

//获取
func GetAll(ctx context.Context) (metadata.MD, bool) {
	return metadata.FromIncomingContext(ctx)
}

func GetVal(ctx context.Context, key string) (string, bool) {
	value, ok := getValue(ctx, key)
	if !ok {
		return "", ok
	}
	return value[0], true
}

func SetVal(ctx context.Context, key string, val string) (context.Context, error) {
	if val == "" {
		return ctx, nil
	}
	pairs := metadata.Pairs(key, val)
	return metadata.NewOutgoingContext(ctx, pairs), nil
}
