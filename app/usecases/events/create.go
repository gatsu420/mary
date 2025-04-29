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
	Name map[string]struct{}
}

func (u *usecaseImpl) CreateEvent(ctx context.Context, arg *CreateEventParams) error {
	events, err := u.cache.GetEvents(ctx, cache.GetEventParams{
		Name: arg.Name,
	})
	if err != nil {
		return errors.New(errors.InternalServerError, "cache failed to get events")
	}

	params := []repository.CreateEventParams{}
	for _, v := range events {
		for _, ev := range v.CreatedAtEpoch {
			params = append(params, repository.CreateEventParams{
				Name:   v.Name,
				UserID: v.UserID,
				CreatedAt: pgtype.Timestamptz{
					Time:  time.Unix(ev, 0),
					Valid: true,
				},
			})
		}
	}

	if _, err = u.query.CreateEvent(ctx, params); err != nil {
		return errors.New(errors.InternalServerError, "DB failed to create event")
	}
	return nil
}
