package workers

import (
	"context"
	"time"

	"github.com/gatsu420/mary/app/usecases/authn"
	"github.com/gatsu420/mary/app/usecases/events"
)

type Worker interface {
	Create(ctx context.Context, ticker <-chan time.Time)
	CreateMembershipRegistry(ctx context.Context)
}

type workerImpl struct {
	authnUsecase authn.Usecase
	usecase      events.Usecase
}

var _ Worker = (*workerImpl)(nil)

func New(authnUsecase authn.Usecase, usecase events.Usecase) Worker {
	return &workerImpl{
		authnUsecase: authnUsecase,
		usecase:      usecase,
	}
}
