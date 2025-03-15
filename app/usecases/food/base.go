package food

import (
	"context"
	"fmt"

	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/db/repository"
	"github.com/rs/zerolog/log"
)

type Usecases interface {
	CreateFood(ctx context.Context, arg *CreateFoodParams) *errors.Error
	ListFood(ctx context.Context, arg *ListFoodParams) ([]repository.ListFoodRow, *errors.Error)
	GetFood(ctx context.Context, id int32) (repository.GetFoodRow, *errors.Error)
	UpdateFood(ctx context.Context, arg *UpdateFoodParams) *errors.Error
	DeleteFood(ctx context.Context, id int32) *errors.Error
	CheckFoodIsRemoved(ctx context.Context, id int32) *errors.Error
}

type usecase struct {
	q *repository.Queries
}

func NewUsecases(q *repository.Queries) Usecases {
	return &usecase{
		q: q,
	}
}

func (u *usecase) CheckFoodIsRemoved(ctx context.Context, id int32) *errors.Error {
	isRemoved, err := u.q.CheckFoodIsRemoved(ctx, id)
	fmt.Println(isRemoved)
	log.Info().Msgf("%v", isRemoved)
	if err != nil {
		return errors.InternalServer("DB failed to check if food is removed")
	}
	if isRemoved {
		return errors.NotFound("food is already removed")
	}

	return nil
}
