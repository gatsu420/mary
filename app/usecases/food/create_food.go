package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

func (u *usecase) CreateFood(ctx context.Context, arg *dbgen.CreateFoodParams) error {
	return u.q.CreateFood(ctx, arg)
}
