package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

func (u *usecase) ListFood(ctx context.Context, arg *dbgen.ListFoodParams) ([]dbgen.ListFoodRow, error) {
	return u.q.ListFood(ctx, arg)
}
