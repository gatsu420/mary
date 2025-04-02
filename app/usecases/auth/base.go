package auth

import (
	"context"

	"github.com/gatsu420/mary/db/repository"
)

type Usecase interface {
	IssueToken(username string) (string, error)
	ValidateToken(signedToken string) (string, error)
	CheckUserIsExisting(ctx context.Context, username string) error
}

type usecaseImpl struct {
	secret string
	query  repository.Querier
}

func NewUsecase(secret string, query repository.Querier) Usecase {
	return &usecaseImpl{
		secret: secret,
		query:  query,
	}
}
