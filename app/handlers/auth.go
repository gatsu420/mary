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

func (as *AuthServer) IssueToken(_ context.Context, user *apiauthv1.IssueTokenRequest) (*apiauthv1.IssueTokenResponse, error) {
	signedToken, err := as.Services.IssueToken(user.UserId)
	if err != nil {
		return nil, err
	}

	return &apiauthv1.IssueTokenResponse{SignedToken: signedToken}, nil
}
