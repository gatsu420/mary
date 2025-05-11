package cache

import (
	"context"
)

type CreateMembershipRegistryParams struct {
	Registry []string
}

func (s *Store) CreateMembershipRegistry(ctx context.Context, arg CreateMembershipRegistryParams) error {
	cmd := s.valkeyClient.B().Rpush().Key("membershipRegistry").
		Element(arg.Registry...).
		Build()
	if err := s.valkeyClient.Do(ctx, cmd).Error(); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetMembershipRegistry(ctx context.Context) ([]string, error) {
	cmd := s.valkeyClient.B().Lrange().Key("membershipRegistry").
		Start(0).Stop(-1).
		Build()
	registry, err := s.valkeyClient.Do(ctx, cmd).AsStrSlice()
	if err != nil {
		return nil, err
	}
	return registry, nil
}
