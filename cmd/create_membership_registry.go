package cmd

import (
	"context"
	"fmt"

	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/authn"
	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/app/workers"
	"github.com/gatsu420/mary/common/config"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/dependency/pgdep"
	"github.com/gatsu420/mary/dependency/valkeydep"
	"github.com/urfave/cli/v2"
)

var CreateMembershipRegistryCmd = &cli.Command{
	Name: "create-membership-registry",
	// Usage: "Create membership registry that should be run only once ",
	Usage: `Create membership registry that enables app to check if a username exists.
	It uses bloom filter stored as list in cache instead of calling DB. This command
	should be run only once to bootstrap the registry.
	`,
	Action: func(ctx *cli.Context) error {
		cfg, err := config.New(".env")
		if err != nil {
			return errors.New(errors.InternalServerError,
				fmt.Sprintf("failed to read config file: %v", err))
		}

		dbPool, err := pgdep.NewPool(cfg.PostgresDSN)
		if err != nil {
			return errors.New(errors.InternalServerError,
				fmt.Sprintf("failed to create DB connection: %v", err))
		}
		defer dbPool.Close()
		dbQuerier := repository.New(dbPool)

		valkeyClient, err := valkeydep.New(cfg.CacheAddress)
		if err != nil {
			return errors.New(errors.InternalServerError,
				fmt.Sprintf("failed to create cache connection: %v", err))
		}
		defer valkeyClient.Close()
		cacheStorer := cache.New(valkeyClient)

		auth := auth.NewAuth(cfg)
		qq := auth.CreateMembershipRegistry([]string{"testUserA", "testUserB"})
		fmt.Println(qq)
		fmt.Println(qq[6], qq[13], qq[14], qq[15])
		fmt.Println(6, 13, 14, 15)

		authnUsecase := authn.NewUsecase(auth, dbQuerier, cacheStorer)
		eventsUsecase := events.NewUsecase(dbQuerier, cacheStorer)

		worker := workers.New(authnUsecase, eventsUsecase)
		workerCtx := context.Background()

		worker.CreateMembershipRegistry(workerCtx)

		return nil
	},
}
