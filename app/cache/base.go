package cache

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

type Storer interface {
	CreateEvent(ctx context.Context, arg CreateEventParams) error
	GetEvents(ctx context.Context, arg GetEventParams) ([]GetEventResponse, error)
}

type Store struct {
	valkeyClient valkey.Client
}

var _ Storer = (*Store)(nil)

func New(valkeyClient valkey.Client) Storer {
	return &Store{
		valkeyClient: valkeyClient,
	}
}
