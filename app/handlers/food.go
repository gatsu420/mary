package handlers

import (
	"context"

	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/common/pbwrappers"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (fs *FoodServer) Create(ctx context.Context, req *apifoodv1.CreateRequest) (*apifoodv1.CreateResponse, error) {
	params := &food.CreateFoodParams{
		Name:           req.Name,
		TypeID:         req.TypeId,
		IntakeStatusID: req.IntakeStatusId,
		FeederID:       req.FeederId,
		LocationID:     req.LocationId,
		Remarks:        pbwrappers.ToPGText(req.Remarks),
	}

	if err := fs.usecase.CreateFood(ctx, params); err != nil {
		return nil, err
	}
	return &apifoodv1.CreateResponse{}, nil
}

func (fs *FoodServer) List(ctx context.Context, req *apifoodv1.ListRequest) (*apifoodv1.ListResponse, error) {
	params := &food.ListFoodParams{
		StartTimestamp: pgtype.Timestamptz{
			Time:  req.StartTimestamp.AsTime(),
			Valid: true,
		},
		EndTimestamp: pgtype.Timestamptz{
			Time:  req.EndTimestamp.AsTime(),
			Valid: true,
		},
		Type:         pbwrappers.ToPGText(req.Type),
		IntakeStatus: pbwrappers.ToPGText(req.IntakeStatus),
		Feeder:       pbwrappers.ToPGText(req.Feeder),
		Location:     pbwrappers.ToPGText(req.Location),
	}
	dbRows, err := fs.usecase.ListFood(ctx, params)
	if err != nil {
		return nil, err
	}

	list := &apifoodv1.ListResponse{}
	for _, r := range dbRows {
		list.Food = append(list.Food, &apifoodv1.ListResponse_Row{
			Id:           r.ID,
			Name:         r.Name,
			Type:         r.Type.String,
			IntakeStatus: r.IntakeStatus.String,
			Feeder:       r.Feeder.String,
			Location:     r.Location.String,
			Remarks:      r.Remarks.String,
			CreatedAt:    timestamppb.New(r.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(r.UpdatedAt.Time),
		})
	}
	return list, nil
}

func (fs *FoodServer) Get(ctx context.Context, req *apifoodv1.GetRequest) (*apifoodv1.GetResponse, error) {
	dbRow, err := fs.usecase.GetFood(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	output := &apifoodv1.GetResponse{
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

func (fs *FoodServer) Update(ctx context.Context, req *apifoodv1.UpdateRequest) (*apifoodv1.UpdateResponse, error) {
	params := &food.UpdateFoodParams{
		Name:           pbwrappers.ToPGText(req.Name),
		TypeID:         pbwrappers.ToPGInt4(req.TypeId),
		IntakeStatusID: pbwrappers.ToPGInt4(req.IntakeStatusId),
		FeederID:       pbwrappers.ToPGInt4(req.FeederId),
		LocationID:     pbwrappers.ToPGInt4(req.LocationId),
		Remarks:        pbwrappers.ToPGText(req.Remarks),
		ID:             req.Id,
	}

	if err := fs.usecase.UpdateFood(ctx, params); err != nil {
		return nil, err
	}
	return &apifoodv1.UpdateResponse{}, nil
}

func (fs *FoodServer) Delete(ctx context.Context, req *apifoodv1.DeleteRequest) (*apifoodv1.DeleteResponse, error) {
	if err := fs.usecase.DeleteFood(ctx, req.Id); err != nil {
		return nil, err
	}
	return &apifoodv1.DeleteResponse{}, nil
}
