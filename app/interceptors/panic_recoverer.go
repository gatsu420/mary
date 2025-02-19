package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func RecoverPanic() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if recv := recover(); recv != nil {
				err = fmt.Errorf("%v", recv)
			}
		}()

		return handler(ctx, req)
	}
}
