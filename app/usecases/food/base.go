package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

type Usecases interface {
	CreateFood(ctx context.Context, arg *dbgen.CreateFoodParams) error
	ListFood(ctx context.Context, arg *dbgen.ListFoodParams) ([]dbgen.ListFoodRow, error)
	GetFood(ctx context.Context, id int32) (dbgen.GetFoodRow, error)
}

type usecase struct {
	q *dbgen.Queries
}

func NewUsecases(q *dbgen.Queries) Usecases {
	return &usecase{
		q: q,
	}
}
