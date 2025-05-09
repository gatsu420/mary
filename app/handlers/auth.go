package handlers

import (
	"context"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
)

func (s *AuthServer) IssueToken(ctx context.Context, user *apiauthv1.IssueTokenRequest) (*apiauthv1.IssueTokenResponse, error) {
	membershipRegistry, err := s.auth.CreateMembershipRegistry()
	if err != nil {
		return nil, err
	}

	if err = s.auth.CheckMembership(membershipRegistry, user.Username); err != nil {
		return nil, err
	}

	signedToken, err := s.auth.IssueToken(user.Username)
	if err != nil {
		return nil, err
	}
	return &apiauthv1.IssueTokenResponse{
		SignedToken: signedToken,
	}, nil
}
