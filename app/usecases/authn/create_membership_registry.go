package authn

import (
	"context"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/common/errors"
)

func (u *usecaseImpl) CreateMembershipRegistry(ctx context.Context) error {
	users, err := u.query.ListUsers(ctx)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to list users")
	}

	registry := u.auth.CreateMembershipRegistry(users)

	params := cache.CreateMembershipRegistryParams{
		Registry: registry,
	}
	if err = u.cache.CreateMembershipRegistry(ctx, params); err != nil {
		return errors.New(errors.InternalServerError, "cache failed to create membership registry")
	}

	return nil
}
