package food

import (
	"context"

	"github.com/gatsu420/mary/app/usecases"
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
	q *repository.Queries
}

func NewUsecases(q *repository.Queries) Usecases {
	return &usecase{
		q: q,
	}
}

func (u *usecase) CheckFoodIsRemoved(ctx context.Context, id int32) error {
	isRemoved, err := u.q.CheckFoodIsRemoved(ctx, id)
	if err != nil {
		return usecases.ErrInternal()
	}
	if isRemoved {
		return usecases.ErrRemovedRecord()
	}

	return nil
}
