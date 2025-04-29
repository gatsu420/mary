package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/common/ctxvalue"
	"github.com/gatsu420/mary/common/tempvalue"
)

type CreateEventParams struct {
	Name string
}

func (s *Store) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	userCtx := ctxvalue.GetUser(ctx)
	listKey := fmt.Sprintf("%v:%v", userCtx.UserID, arg.Name)
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
	Name map[string]struct{}
}

type GetEventResponse struct {
	Name           string
	UserID         string
	CreatedAtEpoch []int64
}

func (s *Store) GetEvents(ctx context.Context, arg GetEventParams) ([]GetEventResponse, error) {
	userID := tempvalue.GetUserID()
	resp := []GetEventResponse{}

	for k := range arg.Name {
		listKey := fmt.Sprintf("%v:%v", userID, k)
		cmd := s.valkeyClient.B().Lrange().Key(listKey).
			Start(0).Stop(-1).
			Build()
		events, err := s.valkeyClient.Do(ctx, cmd).AsIntSlice()
		if err != nil {
			return nil, err
		}

		resp = append(resp, GetEventResponse{
			Name:           k,
			UserID:         userID,
			CreatedAtEpoch: events,
		})
	}

	return resp, nil
}

type DeleteEventParams struct {
	Name map[string]struct{}
}

func (s *Store) DeleteEvents(ctx context.Context, arg DeleteEventParams) error {
	userID := tempvalue.GetUserID()
	for k := range arg.Name {
		listKey := fmt.Sprintf("%v:%v", userID, k)
		cmd := s.valkeyClient.B().Del().Key(listKey).
			Build()
		if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
			return err
		}
	}

	return nil
}
