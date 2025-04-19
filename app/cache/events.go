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
	currentTime := fmt.Sprintf("%v", time.Now().Unix())

	cmd := s.valkeyClient.B().Hset().Key(userCtx.UserID).FieldValue().
		FieldValue(currentTime, arg.Name).
		Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}
