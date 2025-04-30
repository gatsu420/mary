package food

import (
	"context"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/common/tempvalue"
)

func (u *usecaseImpl) DeleteFood(ctx context.Context, id int32) error {
	rows, err := u.query.DeleteFood(ctx, id)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to delete food")
	}
	if rows == 0 {
		return errors.New(errors.NotFoundError, "food not found")
	}

	tempvalue.SetCalledMethods(cache.DeleteFoodEventName)
	eventParams := cache.CreateEventParams{
		Name: cache.DeleteFoodEventName,
	}
	if err := u.cache.CreateEvent(ctx, eventParams); err != nil {
		return errors.New(errors.InternalServerError, "cache failed to create event")
	}

	return nil
}
