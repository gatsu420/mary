package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
)

func (u *usecaseImpl) DeleteFood(ctx context.Context, id int32) error {
	rows, err := u.query.DeleteFood(ctx, id)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to delete food")
	}
	if rows == 0 {
		return errors.New(errors.NotFoundError, "food not found")
	}
	return nil
}
