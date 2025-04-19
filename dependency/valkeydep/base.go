package valkeydep

import "github.com/valkey-io/valkey-go"

func New(address string) (valkey.Client, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{address},
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
