package events

import (
	"context"
	"time"

	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/common/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateEventParams struct {
	Name string
}

func (u *usecaseImpl) CreateEvent(ctx context.Context, arg *CreateEventParams) error {
	events, err := u.cache.GetEvents(ctx, cache.GetEventParams{
		Name: arg.Name,
	})
	if err != nil {
		return errors.New(errors.InternalServerError, "cache failed to get events")
	}

	params := []repository.CreateEventParams{}
	for i := range events.CreatedAtEpoch {
		params = append(params, repository.CreateEventParams{
			Name:   events.Name,
			UserID: events.UserID,
			CreatedAt: pgtype.Timestamptz{
				Time:  time.Unix(events.CreatedAtEpoch[i], 0),
				Valid: true,
			},
		})
	}
	_, err = u.query.CreateEvent(ctx, params)
	if err != nil {
		return errors.New(errors.InternalServerError, "DB failed to create event")
	}
	return nil
}
