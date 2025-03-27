package food

import (
	"context"

	"github.com/gatsu420/mary/db/repository"
)

type Usecases interface {
	CreateFood(ctx context.Context, arg *CreateFoodParams) error
	ListFood(ctx context.Context, arg *ListFoodParams) ([]repository.ListFoodRow, error)
	GetFood(ctx context.Context, id int32) (repository.GetFoodRow, error)
	UpdateFood(ctx context.Context, arg *UpdateFoodParams) error
	DeleteFood(ctx context.Context, id int32) error
}

type usecase struct {
	q *repository.Queries
}

func NewUsecases(q *repository.Queries) Usecases {
	return &usecase{
		q: q,
	}
}
