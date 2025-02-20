package handlers

import (
	"context"
	"fmt"

	"github.com/gatsu420/mary/app/api"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/db/dbgen"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

var foodList *api.FoodList

func init() {
	foodList = &api.FoodList{}
	foodList.Food = []*api.Food{}
}

type FoodServer struct {
	api.UnimplementedFoodServiceServer
	Usecases food.Usecases
}

func (fs *FoodServer) Create(ctx context.Context, food *api.Food) (*emptypb.Empty, error) {
	params := &dbgen.CreateFoodParams{
		Name:           food.Name,
		TypeID:         int32(food.TypeId),
		IntakeStatusID: int32(food.IntakeStatusId),
		FeederID:       int32(food.FeederId),
		LocationID:     int32(food.LocationId),
		Remarks:        pgtype.Text{String: food.Remarks, Valid: food.Remarks != ""},
	}
	if err := fs.Usecases.CreateFood(ctx, params); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (fs *FoodServer) List(_ context.Context, _ *emptypb.Empty) (*api.FoodList, error) {
	log.Info().Msg(fmt.Sprintf("food: %v", foodList.String()))

	return foodList, nil
}
