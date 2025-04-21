package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/common/ctxvalue"
)

type CreateEventParams struct {
	Name string
}

func (s *Store) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	userCtx := ctxvalue.GetUser(ctx)
	listKey := fmt.Sprintf("%v:%v",
		userCtx.UserID,
		arg.Name,
	)
	currentTime := fmt.Sprintf("%v", time.Now().Unix())

	cmd := s.valkeyClient.B().Lpush().Key(listKey).
		Element(currentTime).
		Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}

type GetEventParams struct {
	Name string
}

type GetEventResponse struct {
	Name           string
	UserID         string
	CreatedAtEpoch []int64
}

func (s *Store) GetEvents(ctx context.Context, arg GetEventParams) (GetEventResponse, error) {
	userCtx := ctxvalue.GetUser(ctx)
	listKey := fmt.Sprintf("%v:%v",
		userCtx.UserID,
		arg.Name,
	)

	cmd := s.valkeyClient.B().Lrange().Key(listKey).
		Start(0).Stop(-1).
		Build()
	events, err := s.valkeyClient.Do(ctx, cmd).AsIntSlice()
	if err != nil {
		return GetEventResponse{}, err
	}
	return GetEventResponse{
		Name:           arg.Name,
		UserID:         userCtx.UserID,
		CreatedAtEpoch: events,
	}, nil
}
