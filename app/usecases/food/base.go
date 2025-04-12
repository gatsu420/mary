package food

import (
	"context"

	"github.com/gatsu420/mary/db/repository"
)

type Usecase interface {
	CreateFood(ctx context.Context, arg *CreateFoodParams) error
	ListFood(ctx context.Context, arg *ListFoodParams) ([]repository.ListFoodRow, error)
	GetFood(ctx context.Context, id int32) (repository.GetFoodRow, error)
	UpdateFood(ctx context.Context, arg *UpdateFoodParams) error
	DeleteFood(ctx context.Context, id int32) error
}

type usecaseImpl struct {
	query repository.Querier
}

func NewUsecase(query repository.Querier) Usecase {
	return &usecaseImpl{
		query: query,
	}
}
