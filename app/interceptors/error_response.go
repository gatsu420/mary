package interceptors

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ResponseError() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			if svcerr, ok := err.(*errors.Error); ok {
				return res, status.Error(svcerr.GRPCCode(), svcerr.Error())
			}
		}

		return res, nil
	}
}
