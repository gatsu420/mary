package interceptors

import (
	"context"

	"github.com/gatsu420/mary/app/usecases/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey int

const authTokenClaimCtx ctxKey = iota

func ValidateToken(authUsecases auth.Usecases) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		publicMethods := map[string]bool{
			"/auth.v1.AuthService/IssueToken": true,
		}
		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		signedToken := md["authorization"][0]
		if len(signedToken) == 0 && !publicMethods[info.FullMethod] {
			return nil, status.Error(codes.Unauthenticated, "token is not provided")
		}

		userID, err := authUsecases.ValidateToken(signedToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "%v", err)
		}

		ctx = context.WithValue(ctx, authTokenClaimCtx, userID)

		return handler(ctx, req)
	}
}
