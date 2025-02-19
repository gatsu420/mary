package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Auth interface {
	ValidateToken(signedToken string) (string, error)
}

type authInterceptor struct {
	auth Auth
}

func NewAuthInterceptor(auth Auth) (*authInterceptor, error) {
	if auth == nil {
		return nil, fmt.Errorf("cannot have nil auth")
	}

	return &authInterceptor{auth}, nil
}

func (i *authInterceptor) ValidateTokenUnaryInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	publicMethods := map[string]bool{
		"/api.AuthService/IssueToken": true,
	}
	if publicMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	token := md["authorization"][0]
	if len(token) == 0 && !publicMethods[info.FullMethod] {
		return nil, status.Error(codes.Unauthenticated, "JWT is not provided")
	}

	userID, err := i.auth.ValidateToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	ctx = context.WithValue(ctx, "userID", userID)

	return handler(ctx, req)
}
