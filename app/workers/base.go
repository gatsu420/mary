package workers

import (
	"context"
	"time"

	"github.com/gatsu420/mary/app/usecases/events"
)

type Worker interface {
	Create(ctx context.Context, ticker <-chan time.Time)
}

type workerImpl struct {
	usecase events.Usecase
}

var _ Worker = (*workerImpl)(nil)

func New(usecase events.Usecase) Worker {
	return &workerImpl{
		usecase: usecase,
	}
}
