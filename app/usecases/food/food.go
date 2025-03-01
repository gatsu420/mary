package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
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
	params := &dbgen.CreateFoodParams{
		Name:           arg.Name,
		TypeID:         arg.TypeID,
		IntakeStatusID: arg.IntakeStatusID,
		FeederID:       arg.FeederID,
		LocationID:     arg.LocationID,
		Remarks:        arg.Remarks,
	}

	return u.q.CreateFood(ctx, params)
}

type ListFoodParams struct {
	StartTimestamp pgtype.Timestamptz
	EndTimestamp   pgtype.Timestamptz
	Type           pgtype.Text
	IntakeStatus   pgtype.Text
	Feeder         pgtype.Text
	Location       pgtype.Text
}

func (u *usecase) ListFood(ctx context.Context, arg *ListFoodParams) ([]dbgen.ListFoodRow, error) {
	params := &dbgen.ListFoodParams{
		StartTimestamp: arg.StartTimestamp,
		EndTimestamp:   arg.EndTimestamp,
		Type:           arg.Type,
		IntakeStatus:   arg.IntakeStatus,
		Feeder:         arg.Feeder,
		Location:       arg.Location,
	}

	return u.q.ListFood(ctx, params)
}

func (u *usecase) GetFood(ctx context.Context, id int32) (dbgen.GetFoodRow, error) {
	return u.q.GetFood(ctx, id)
}
