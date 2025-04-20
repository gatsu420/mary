package food

import (
	"context"
	stderr "errors"

	"github.com/gatsu420/mary/app/cache"
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

	if err != nil {
		if stderr.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(errors.NotFoundError, "food not found")
		}
		return nil, errors.New(errors.InternalServerError, "DB failed to get food")
	}

	eventParams := cache.CreateEventParams{
		Name: "GetFood",
	}
	if err := u.cache.CreateEvent(ctx, eventParams); err != nil {
		return nil, errors.New(errors.InternalServerError, "cache failed to create event")
	}

	return row, nil
}
