package helper

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func MetadataValue(ctx context.Context, key string) ([]string, bool) {
	incomingContext, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ok
	}
	strings, ok := incomingContext[key]
	return strings, ok
}
func MetadataAll(ctx context.Context) (metadata.MD, bool) {
	return metadata.FromIncomingContext(ctx)

}
