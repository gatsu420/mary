package interceptors

import (
	"context"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/common/ctxvalue"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/common/tempvalue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ValidateToken(auth auth.Auth) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		publicMethods := map[string]bool{
			"/auth.v1.AuthService/IssueToken": true,
		}
		if publicMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, _ := metadata.FromIncomingContext(ctx)
		signedToken := md["authorization"][0]
		if len(signedToken) == 0 {
			return nil, errors.New(errors.InternalServerError, "token is not provided")
		}

		userID, err := auth.ValidateToken(signedToken)
		if err != nil {
			return nil, err
		}

		tempvalue.SetUserID(userID)

		userCtx := ctxvalue.User{
			UserID: userID,
		}
		ctx = ctxvalue.SetUser(ctx, userCtx)
		return handler(ctx, req)
	}
}
