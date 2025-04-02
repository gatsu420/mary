package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/jackc/pgx/v5"
)

func (u *usecaseImpl) GetFood(ctx context.Context, id int32) (repository.GetFoodRow, error) {
	food, err := u.query.GetFood(ctx, id)
	switch err {
	case pgx.ErrNoRows:
		return repository.GetFoodRow{}, errors.New(errors.NotFoundError, "food not found")
	case nil:
		return food, nil
	default:
		return repository.GetFoodRow{}, errors.New(errors.InternalServerError, "DB failed to get food")
	}
}
