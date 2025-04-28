package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/gatsu420/mary/app/usecases/events"
)

func (w *workerImpl) Create() {
	for {
		params := &events.CreateEventParams{
			Name: "GetFood",
		}
		if err := w.usecase.CreateEvent(context.Background(), params); err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
	}
}
