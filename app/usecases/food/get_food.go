package food

import (
	"context"

	"github.com/gatsu420/mary/db/dbgen"
)

func (u *usecase) GetFood(ctx context.Context, id int32) (dbgen.GetFoodRow, error) {
	return u.q.GetFood(ctx, id)
}
