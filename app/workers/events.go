package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/common/tempvalue"
)

func (w *workerImpl) Create(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
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

		time.Sleep(1 * time.Minute)
	}
}
