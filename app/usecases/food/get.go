package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetFoodRow struct {
	ID           int32
	Name         string
	Type         pgtype.Text
	IntakeStatus pgtype.Text
	Feeder       pgtype.Text
	Location     pgtype.Text
	Remarks      pgtype.Text
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

func (u *usecaseImpl) GetFood(ctx context.Context, id int32) (*GetFoodRow, error) {
	food, err := u.query.GetFood(ctx, id)

	row := &GetFoodRow{
		ID:           food.ID,
		Name:         food.Name,
		Type:         food.Type,
		IntakeStatus: food.IntakeStatus,
		Feeder:       food.Feeder,
		Location:     food.Location,
		Remarks:      food.Remarks,
		CreatedAt:    food.CreatedAt,
		UpdatedAt:    food.UpdatedAt,
	}

	switch err {
	case pgx.ErrNoRows:
		return &GetFoodRow{}, errors.New(errors.NotFoundError, "food not found")
	case nil:
		return row, nil
	default:
		return &GetFoodRow{}, errors.New(errors.InternalServerError, "DB failed to get food")
	}
}
