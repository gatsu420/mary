package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/jackc/pgx/v5"
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

func (u *usecase) GetFood(ctx context.Context, id int32) (repository.GetFoodRow, error) {
	// TODO: move this not in the same layer as usecase, maybe in interceptor?
	// also for update and delete method that requires checking food removal status
	// if err := u.CheckFoodIsRemoved(ctx, id); err != nil {
	// 	return repository.GetFoodRow{}, err
	// }

	food, err := u.q.GetFood(ctx, id)
	switch err {
	case pgx.ErrNoRows:
		return repository.GetFoodRow{}, errors.New(errors.NotFoundError, "food not found")
	case nil:
		return food, nil
	default:
		return repository.GetFoodRow{}, errors.New(errors.InternalServerError, "DB failed to get food")
	}
}

type UpdateFoodParams struct {
	Name           pgtype.Text
	TypeID         pgtype.Int4
	IntakeStatusID pgtype.Int4
	FeederID       pgtype.Int4
	LocationID     pgtype.Int4
	Remarks        pgtype.Text
	ID             int32
}

func (u *usecase) UpdateFood(ctx context.Context, arg *UpdateFoodParams) error {
	// if err := u.CheckFoodIsRemoved(ctx, arg.ID); err != nil {
	// 	return err
	// }

	params := &repository.UpdateFoodParams{
		Name:           arg.Name,
		TypeID:         arg.TypeID,
		IntakeStatusID: arg.IntakeStatusID,
		FeederID:       arg.FeederID,
		LocationID:     arg.LocationID,
		Remarks:        arg.Remarks,
		ID:             arg.ID,
	}

	if err := u.q.UpdateFood(ctx, params); err != nil {
		return errors.New(errors.InternalServerError, "DB failed to update food")
	}
	return nil
}

func (u *usecase) DeleteFood(ctx context.Context, id int32) error {
	// if err := u.CheckFoodIsRemoved(ctx, id); err != nil {
	// 	return err
	// }

	if err := u.q.DeleteFood(ctx, id); err != nil {
		return errors.New(errors.InternalServerError, "DB failed to delete food")
	}
	return nil
}
