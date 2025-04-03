package handlers

import (
	"context"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	"github.com/gatsu420/mary/app/usecases/users"
	"github.com/gatsu420/mary/auth"
)

type AuthServer struct {
	apiauthv1.UnimplementedAuthServiceServer
	Auth         auth.Auth
	UsersUsecase users.Usecase
}

func (s *AuthServer) IssueToken(ctx context.Context, user *apiauthv1.IssueTokenRequest) (*apiauthv1.IssueTokenResponse, error) {
	if err := s.UsersUsecase.CheckUserIsExisting(ctx, user.Username); err != nil {
		return nil, err
	}

	signedToken, err := s.Auth.IssueToken(user.Username)
	if err != nil {
		return nil, err
	}
	return &apiauthv1.IssueTokenResponse{
		SignedToken: signedToken,
	}, nil
}
