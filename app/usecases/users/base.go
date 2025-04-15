package users

import (
	"context"

	"github.com/gatsu420/mary/app/repository"
)

type Usecase interface {
	CheckUserIsExisting(ctx context.Context, username string) error
}

type usecaseImpl struct {
	query repository.Querier
}

var _ Usecase = (*usecaseImpl)(nil)

func NewUsecase(query repository.Querier) Usecase {
	return &usecaseImpl{
		query: query,
	}
}
