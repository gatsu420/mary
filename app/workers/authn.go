package workers

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (w *workerImpl) CreateMembershipRegistry(ctx context.Context) {
	if err := w.authnUsecase.CreateMembershipRegistry(ctx); err != nil {
		log.Fatal().Msgf("worker failed to create membership registry: %v", err)
	}
}
