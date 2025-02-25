package handlers

import (
	"context"

	"github.com/gatsu420/mary/app/api"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/db/dbgen"
	"github.com/gatsu420/mary/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FoodServer struct {
	api.UnimplementedFoodServiceServer
	Usecases food.Usecases
}

func (fs *FoodServer) Create(ctx context.Context, food *api.CreateRequest) (*api.CreateResponse, error) {
	params := &dbgen.CreateFoodParams{
		Name:           food.Name,
		TypeID:         food.TypeId,
		IntakeStatusID: food.IntakeStatusId,
		FeederID:       food.FeederId,
		LocationID:     food.LocationId,
		Remarks:        pgtype.Text{String: food.Remarks, Valid: food.Remarks != ""},
	}
	if err := fs.Usecases.CreateFood(ctx, params); err != nil {
		return nil, err
	}

	return &api.CreateResponse{}, nil
}

func (fs *FoodServer) List(ctx context.Context, req *api.ListRequest) (resp *api.ListResponse, err error) {
	params := &dbgen.ListFoodParams{
		StartTimestamp: pgtype.Timestamptz{
			Time:  req.StartTimestamp.AsTime(),
			Valid: true,
		},
		EndTimestamp: pgtype.Timestamptz{
			Time:  req.EndTimestamp.AsTime(),
			Valid: true,
		},
		Type:         utils.NullStringWrapperToPGText(req.Type),
		IntakeStatus: utils.NullStringWrapperToPGText(req.IntakeStatus),
		Feeder:       utils.NullStringWrapperToPGText(req.Feeder),
		Location:     utils.NullStringWrapperToPGText(req.Location),
	}
	dbRows, err := fs.Usecases.ListFood(ctx, params)
	if err != nil {
		return nil, err
	}

	list := &api.ListResponse{}
	for _, r := range dbRows {
		list.Food = append(list.Food, &api.ListResponse_Row{
			Id:           r.ID,
			Name:         r.Name,
			Type:         r.Type.String,
			IntakeStatus: r.IntakeStatus.String,
			Feeder:       r.Feeder.String,
			Location:     r.Location.String,
			CreatedAt:    timestamppb.New(r.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(r.UpdatedAt.Time),
		})
	}

	return list, nil
}

func (fs *FoodServer) Get(ctx context.Context, req *api.GetRequest) (resp *api.GetResponse, err error) {
	dbRow, err := fs.Usecases.GetFood(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	output := &api.GetResponse{
		Id:           dbRow.ID,
		Name:         dbRow.Name,
		Type:         dbRow.Type.String,
		IntakeStatus: dbRow.IntakeStatus.String,
		Feeder:       dbRow.Feeder.String,
		Location:     dbRow.Location.String,
		Remarks:      dbRow.Remarks.String,
		CreatedAt:    timestamppb.New(dbRow.CreatedAt.Time),
		UpdatedAt:    timestamppb.New(dbRow.UpdatedAt.Time),
	}

	return output, nil
}
