package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/common/tempvalue"
)

func (w *workerImpl) Create() {
	for {
		params := &events.CreateEventParams{
			Name: tempvalue.GetCalledMethods(),
		}
		if err := w.usecase.CreateEvent(context.Background(), params); err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Minute)
	}
}
