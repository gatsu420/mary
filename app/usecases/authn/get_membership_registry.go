package authn

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
)

func (u *usecaseImpl) GetMembershipRegistry(ctx context.Context) ([]string, error) {
	registry, err := u.cache.GetMembershipRegistry(ctx)
	if err != nil {
		return nil, errors.New(errors.InternalServerError, "cache failed to get membership registry")
	}
	return registry, nil
}
