package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/common/tempvalue"
)

func (w *workerImpl) Create(ctx context.Context, ticker <-chan time.Time) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker:
			calledMethods := tempvalue.GetCalledMethods()
			if len(calledMethods) != 0 {
				params := &events.CreateEventParams{
					Name: calledMethods,
				}
				if err := w.usecase.CreateEvent(context.Background(), params); err != nil {
					fmt.Println(err)
				}
				tempvalue.FlushCalledMethods()
			}
		}
	}
}
