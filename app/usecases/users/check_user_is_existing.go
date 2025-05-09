package users

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
)

// User checking via repository has been deprecated. Use CheckMembership in auth layer instead.
func (u *usecaseImpl) CheckUserIsExisting(ctx context.Context, username string) error {
	isExisting, err := u.query.CheckUserIsExisting(ctx, username)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to check if user is existing")
	}
	if !isExisting {
		return errors.New(errors.NotFoundError, "user not found")
	}
	return nil
}
