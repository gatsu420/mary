package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

type Usecases interface {
	CreateFood(ctx context.Context, arg *dbgen.CreateFoodParams) error
}

type usecase struct {
	q *dbgen.Queries
}

func NewUsecases(q *dbgen.Queries) Usecases {
	return &usecase{
		q: q,
	}
}
