package cache

import (
	"context"
)

type CreateMembershipRegistryParams struct {
	Registry []string
}

func (s *Store) CreateMembershipRegistry(ctx context.Context, arg CreateMembershipRegistryParams) error {
	cmd := s.valkeyClient.B().Lpush().Key("membershipRegistry").
		Element(arg.Registry...).
		Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}
