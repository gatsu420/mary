package food

import (
	"context"
	"fmt"

	"github.com/gatsu420/mary/app/repository"
)

func (u *usecaseImpl) CreateEvent(ctx context.Context) error {
	fmt.Println("usecase is called")

	params := []*repository.CreateEventParams{}
	i := 0
	for i < 1000 {
		params = append(params, &repository.CreateEventParams{
			Name:   "tess",
			UserID: fmt.Sprintf("%v", i),
		})
		i++
	}

	_, err := u.query.CreateEvent(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
