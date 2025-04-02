package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateFoodParams struct {
	Name           pgtype.Text
	TypeID         pgtype.Int4
	IntakeStatusID pgtype.Int4
	FeederID       pgtype.Int4
	LocationID     pgtype.Int4
	Remarks        pgtype.Text
	ID             int32
}

func (u *usecaseImpl) UpdateFood(ctx context.Context, arg *UpdateFoodParams) error {
	params := &repository.UpdateFoodParams{
		Name:           arg.Name,
		TypeID:         arg.TypeID,
		IntakeStatusID: arg.IntakeStatusID,
		FeederID:       arg.FeederID,
		LocationID:     arg.LocationID,
		Remarks:        arg.Remarks,
		ID:             arg.ID,
	}

	rows, err := u.query.UpdateFood(ctx, params)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to update food")
	}
	if rows == 0 {
		return errors.New(errors.NotFoundError, "food not found")
	}
	return nil
}
