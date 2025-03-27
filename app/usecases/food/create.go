package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
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

func (u *usecase) CreateFood(ctx context.Context, arg *CreateFoodParams) error {
	params := &repository.CreateFoodParams{
		Name:           arg.Name,
		TypeID:         arg.TypeID,
		IntakeStatusID: arg.IntakeStatusID,
		FeederID:       arg.FeederID,
		LocationID:     arg.LocationID,
		Remarks:        arg.Remarks,
	}

	if err := u.q.CreateFood(ctx, params); err != nil {
		return errors.New(errors.InternalServerError, "DB failed to create food")
	}
	return nil
}
