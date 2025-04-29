package food

import (
	"context"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/common/tempvalue"
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

type ListFoodRow struct {
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

func (u *usecaseImpl) ListFood(ctx context.Context, arg *ListFoodParams) ([]ListFoodRow, error) {
	params := repository.ListFoodParams{
		StartTimestamp: arg.StartTimestamp,
		EndTimestamp:   arg.EndTimestamp,
		Type:           arg.Type,
		IntakeStatus:   arg.IntakeStatus,
		Feeder:         arg.Feeder,
		Location:       arg.Location,
	}

	foodList, err := u.query.ListFood(ctx, params)
	if err != nil {
		return nil, errors.New(errors.InternalServerError, "DB failed to list food")
	}

	rows := []ListFoodRow{}
	for _, f := range foodList {
		rows = append(rows, ListFoodRow{
			ID:           f.ID,
			Name:         f.Name,
			Type:         f.Type,
			IntakeStatus: f.IntakeStatus,
			Feeder:       f.Feeder,
			Location:     f.Location,
			Remarks:      f.Remarks,
			CreatedAt:    f.CreatedAt,
			UpdatedAt:    f.UpdatedAt,
		})
	}

	tempvalue.SetCalledMethods("ListFood")

	eventParams := cache.CreateEventParams{
		Name: "ListFood",
	}
	if err := u.cache.CreateEvent(ctx, eventParams); err != nil {
		return nil, errors.New(errors.InternalServerError, "cache failed to create event")
	}

	return rows, nil
}
