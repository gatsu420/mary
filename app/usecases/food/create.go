package food

import (
	"context"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/common/ctxvalue"
	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateFoodParams struct {
	Name           string
	TypeID         int32
	IntakeStatusID int32
	FeederID       int32
	LocationID     int32
	Remarks        pgtype.Text
}

func (u *usecaseImpl) CreateFood(ctx context.Context, arg *CreateFoodParams) error {
	params := repository.CreateFoodParams{
		Name:           arg.Name,
		TypeID:         arg.TypeID,
		IntakeStatusID: arg.IntakeStatusID,
		FeederID:       arg.FeederID,
		LocationID:     arg.LocationID,
		Remarks:        arg.Remarks,
	}

	if err := u.query.CreateFood(ctx, params); err != nil {
		return errors.New(errors.InternalServerError, "DB failed to create food")
	}

	userCtx := ctxvalue.GetUser(ctx)
	eventParams := cache.CreateEventParams{
		Name:   "CreateFood",
		UserID: userCtx.UserID,
	}
	if err := u.cache.CreateEvent(ctx, eventParams); err != nil {
		return errors.New(errors.InternalServerError, "cache failed to store event")
	}

	return nil
}
