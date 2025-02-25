package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

func (u *usecase) CreateFood(ctx context.Context, arg *dbgen.CreateFoodParams) error {
	return u.q.CreateFood(ctx, arg)
}

func (u *usecase) ListFood(ctx context.Context, arg *dbgen.ListFoodParams) ([]dbgen.ListFoodRow, error) {
	return u.q.ListFood(ctx, arg)
}

func (u *usecase) GetFood(ctx context.Context, id int32) (dbgen.GetFoodRow, error) {
	return u.q.GetFood(ctx, id)
}
