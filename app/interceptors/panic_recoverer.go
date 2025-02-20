package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoverPanic() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if recv := recover(); recv != nil {
				err = status.Errorf(codes.Internal, "internal server error: %v", recv)
			}
		}()

		return handler(ctx, req)
	}
}
