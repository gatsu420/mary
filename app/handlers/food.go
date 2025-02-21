package handlers

import (
	"context"

	"github.com/gatsu420/mary/app/api"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/db/dbgen"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// var foodList *api.FoodList

// func init() {
// 	foodList = &api.FoodList{}
// 	foodList.Food = []*api.Food{}
// }

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

func (fs *FoodServer) List(ctx context.Context, req *api.ListRequest) (resp *api.ListResponse, err error) {
	// log.Info().Msg(fmt.Sprintf("food: %v", foodList.String()))

	// return foodList, nil

	params := &dbgen.ListFoodParams{
		StartTimestamp: pgtype.Timestamptz{
			Time:  req.StartTimestamp.AsTime(),
			Valid: true,
		},
		EndTimestamp: pgtype.Timestamptz{
			Time:  req.EndTimestamp.AsTime(),
			Valid: true,
		},
	}
	dbRows, err := fs.Usecases.ListFood(ctx, params)
	if err != nil {
		return nil, err
	}

	list := &api.ListResponse{}
	for _, r := range dbRows {
		list.FoodList = append(list.FoodList, &api.ListResponseRow{
			Id: r.ID,
			Food: &api.Food{
				Name:           r.Name,
				TypeId:         api.Food_Type(r.TypeID),
				IntakeStatusId: api.Food_IntakeStatus(r.IntakeStatusID),
				FeederId:       api.Food_Feeder(r.FeederID),
				LocationId:     api.Food_Location(r.LocationID),
				Remarks:        r.Remarks.String,
			},
			CreatedAt: timestamppb.New(r.CreatedAt.Time),
			UpdatedAt: timestamppb.New(r.UpdatedAt.Time),
		})
	}

	return list, nil
}
