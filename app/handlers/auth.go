package handlers

import (
	"context"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	"github.com/gatsu420/mary/app/auth"
)

type AuthServer struct {
	apiauthv1.UnimplementedAuthServiceServer
	Services auth.Services
}

func (as *AuthServer) IssueToken(ctx context.Context, user *apiauthv1.IssueTokenRequest) (*apiauthv1.IssueTokenResponse, error) {
	if err := as.Services.CheckUserIsExisting(ctx, user.Username); err != nil {
		return nil, err
	}

	signedToken, err := as.Services.IssueToken(user.Username)
	if err != nil {
		return nil, err
	}

	return &apiauthv1.IssueTokenResponse{SignedToken: signedToken}, nil
}
