package handlers

import (
	"context"

	"github.com/gatsu420/mary/app/api"
	"github.com/gatsu420/mary/app/auth"
)

type AuthServer struct {
	api.UnimplementedAuthServiceServer
	Services auth.Services
}

func (as *AuthServer) IssueToken(_ context.Context, user *api.User) (*api.Token, error) {
	signedToken, err := as.Services.IssueToken(user.UserID)
	if err != nil {
		return nil, err
	}

	return &api.Token{SignedToken: signedToken}, nil
}
