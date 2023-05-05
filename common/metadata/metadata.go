package metadata

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// 获取
func getValue(ctx context.Context, key string) ([]string, bool) {
	incomingContext, ok := GetAll(ctx)
	if !ok {
		return nil, ok
	}
	strings, ok := incomingContext[key]
	return strings, ok
}

// 获取
func GetAll(ctx context.Context) (metadata.MD, bool) {
	return metadata.FromIncomingContext(ctx)
}

func GetVal(ctx context.Context, key string) (string, bool) {
	key = strings.ToLower(key)
	md, ok := GetAll(ctx)
	if !ok {
		return "", ok
	}
	if len(md.Get(key)) == 0 {
		return "", false
	}
	return md.Get(key)[0], ok
}

func GetValUnmarshal(ctx context.Context, key string, dest proto.Message) error {
	val, ok := GetVal(ctx, key)
	if !ok {
		return errors.New("no data")
	}
	return proto.Unmarshal([]byte(val), dest)
}

func SetVal(ctx context.Context, key string, val string) (context.Context, error) {

	if len(val) == 0 {
		return ctx, nil
	}

	return metadata.AppendToOutgoingContext(ctx, strings.ToLower(key), val), nil
}

func GetMessage(ctx context.Context, key string, dest proto.Message) (error, bool) {
	md, ok := GetVal(ctx, key)
	if !ok {
		return nil, ok
	}
	if len(md) == 0 {
		return nil, false
	}
	err := prototext.Unmarshal([]byte(md), dest)
	if err != nil {
		return err, false
	}
	return nil, dest == nil || dest.ProtoReflect().IsValid()
}

func SetMessage(ctx context.Context, key string, val proto.Message) (context.Context, error) {

	if val == nil || !val.ProtoReflect().IsValid() {
		return ctx, nil
	}
	data, err := prototext.Marshal(val)
	if err != nil {
		return ctx, err
	}
	return metadata.AppendToOutgoingContext(ctx, strings.ToLower(key), string(data)), nil
}
