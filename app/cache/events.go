package cache

import "context"

type CreateEventParams struct {
	Name           string
	UserID         string
	CreatedAtEpoch string
}

func (s *Store) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	cmd := s.valkeyClient.B().Hset().Key(arg.UserID).
		FieldValue().FieldValue(arg.CreatedAtEpoch, arg.Name).Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}
