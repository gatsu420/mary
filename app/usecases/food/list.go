package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListFoodParams struct {
	StartTimestamp pgtype.Timestamptz
	EndTimestamp   pgtype.Timestamptz
	Type           pgtype.Text
	IntakeStatus   pgtype.Text
	Feeder         pgtype.Text
	Location       pgtype.Text
}

func (u *usecase) ListFood(ctx context.Context, arg *ListFoodParams) ([]repository.ListFoodRow, error) {
	params := &repository.ListFoodParams{
		StartTimestamp: arg.StartTimestamp,
		EndTimestamp:   arg.EndTimestamp,
		Type:           arg.Type,
		IntakeStatus:   arg.IntakeStatus,
		Feeder:         arg.Feeder,
		Location:       arg.Location,
	}

	foodList, err := u.q.ListFood(ctx, params)
	if err != nil {
		return nil, errors.New(errors.InternalServerError, "DB failed to list food")
	}
	return foodList, nil
}
