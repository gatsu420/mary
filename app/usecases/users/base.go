package users

import (
	"context"

	"github.com/gatsu420/mary/db/repository"
)

type Usecase interface {
	CheckUserIsExisting(ctx context.Context, username string) error
}

type usecaseImpl struct {
	query repository.Querier
}

func NewUsecase(query repository.Querier) Usecase {
	return &usecaseImpl{
		query: query,
	}
}
