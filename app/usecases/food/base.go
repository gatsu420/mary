package food

import (
	"context"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
)

type Usecases interface {
	CreateFood(ctx context.Context, arg *CreateFoodParams) error
	ListFood(ctx context.Context, arg *ListFoodParams) ([]repository.ListFoodRow, error)
	GetFood(ctx context.Context, id int32) (repository.GetFoodRow, error)
	UpdateFood(ctx context.Context, arg *UpdateFoodParams) error
	DeleteFood(ctx context.Context, id int32) error
	CheckFoodIsRemoved(ctx context.Context, id int32) error
}

type usecase struct {
	q repository.Querier
}

func NewUsecases(q repository.Querier) Usecases {
	return &usecase{
		q: q,
	}
}

func (u *usecase) CheckFoodIsRemoved(ctx context.Context, id int32) error {
	isRemoved, err := u.q.CheckFoodIsRemoved(ctx, id)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to check if food is removed")
	}
	if isRemoved {
		return errors.New(errors.NotFoundError, "food not found")
	}

	return nil
}
