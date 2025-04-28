package workers

import "github.com/gatsu420/mary/app/usecases/events"

type Worker interface {
	Start()
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

func (w *workerImpl) Start() {
	go w.Create()
}
