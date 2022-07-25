package metadata

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

//获取
func getValue(ctx context.Context, key string) ([]string, bool) {
	incomingContext, ok := GetAll(ctx)
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
	key = strings.ToLower(key)
	incomingContext, ok := GetAll(ctx)
	if !ok {
		return "", ok
	}
	strings, ok := incomingContext[key]
	if !ok {
		return "", ok
	}
	return strings[0], ok
}

func GetValUnmarshal(ctx context.Context, key string, dest proto.Message) error {
	val, ok := GetVal(ctx, key)
	if !ok {
		return errors.New("no data")
	}
	return proto.Unmarshal([]byte(val), dest)
}

func SetVal(ctx context.Context, key string, val string) (context.Context, error) {
	if val == "" {
		return ctx, nil
	}
	// pairs := metadata.Pairs(key, val)

	return metadata.AppendToOutgoingContext(ctx, strings.ToLower(key), val), nil
}
