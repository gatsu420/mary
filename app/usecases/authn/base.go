package authn

import (
	"context"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
)

type Usecase interface {
	CreateMembershipRegistry(ctx context.Context) error
	GetMembershipRegistry(ctx context.Context) ([]string, error)
}

type usecaseImpl struct {
	auth  auth.Auth
	query repository.Querier
	cache cache.Storer
}

func NewUsecase(auth auth.Auth, query repository.Querier, cache cache.Storer) Usecase {
	return &usecaseImpl{
		auth:  auth,
		query: query,
		cache: cache,
	}
}
