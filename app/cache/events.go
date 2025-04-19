package cache

import (
	"context"
	"fmt"
	"time"
)

type CreateEventParams struct {
	Name   string
	UserID string
}

func (s *Store) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	currentTime := fmt.Sprintf("%v", time.Now().Unix())
	cmd := s.valkeyClient.B().Hset().Key(arg.UserID).FieldValue().
		FieldValue(currentTime, arg.Name).
		Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}
