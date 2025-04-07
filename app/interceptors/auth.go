package interceptors

import (
	"context"

	"github.com/gatsu420/mary/auth"
	"github.com/gatsu420/mary/common/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKey int

const authTokenClaimCtx ctxKey = iota

func ValidateToken(auth auth.Auth) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		publicMethods := map[string]bool{
			"/auth.v1.AuthService/IssueToken": true,
		}
		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New(errors.InternalServerError, "metadata is not provided")
		}

		signedToken := md["authorization"][0]
		if len(signedToken) == 0 && !publicMethods[info.FullMethod] {
			return nil, errors.New(errors.InternalServerError, "token is not provided")
		}

		userID, err := auth.ValidateToken(signedToken)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, authTokenClaimCtx, userID)

		return handler(ctx, req)
	}
}
