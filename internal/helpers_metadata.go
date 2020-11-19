package internal

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const (
	ConstRequestID = "x-request-id"
)

// GetMetadataField...
func GetMetadataField(ctx context.Context, key string) string {
	md, _ := metadata.FromIncomingContext(ctx)
	if field, ok := md[key]; !ok {
		return ""
	} else {
		return field[0]
	}
}
