package events

import (
	"context"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
)

type Usecase interface {
	CreateEvent(ctx context.Context, arg *CreateEventParams) error
}

type usecaseImpl struct {
	query repository.Querier
	cache cache.Storer
}

var _ Usecase = (*usecaseImpl)(nil)

func NewUsecase(query repository.Querier, cache cache.Storer) Usecase {
	return &usecaseImpl{
		query: query,
		cache: cache,
	}
}
