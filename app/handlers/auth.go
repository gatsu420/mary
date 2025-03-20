package handlers

import (
	"context"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	"github.com/gatsu420/mary/app/usecases/auth"
)

type AuthServer struct {
	apiauthv1.UnimplementedAuthServiceServer
	Usecases auth.Usecases
}

func (s *AuthServer) IssueToken(ctx context.Context, user *apiauthv1.IssueTokenRequest) (*apiauthv1.IssueTokenResponse, error) {
	if err := s.Usecases.CheckUserIsExisting(ctx, user.Username); err != nil {
		return nil, err
	}

	signedToken, err := s.Usecases.IssueToken(user.Username)
	if err != nil {
		return nil, err
	}
	return &apiauthv1.IssueTokenResponse{
		SignedToken: signedToken,
	}, nil
}
