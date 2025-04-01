package auth

import (
	"context"

	"github.com/gatsu420/mary/db/repository"
)

type Usecases interface {
	IssueToken(username string) (string, error)
	ValidateToken(signedToken string) (string, error)
	CheckUserIsExisting(ctx context.Context, username string) error
}

type usecase struct {
	secret string
	query  repository.Querier
}

func NewUsecases(secret string, query repository.Querier) Usecases {
	return &usecase{
		secret: secret,
		query:  query,
	}
}
