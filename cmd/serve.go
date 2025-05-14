package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/authn"
	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/app/usecases/users"
	"github.com/gatsu420/mary/app/workers"
	"github.com/gatsu420/mary/common/config"
	"github.com/gatsu420/mary/common/errors"
	"github.com/gatsu420/mary/dependency/pgdep"
	"github.com/gatsu420/mary/dependency/valkeydep"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ServeCmd = &cli.Command{
	Name:  "serve",
	Usage: "Serve API",
	Action: func(ctx *cli.Context) error {
		// Filepath for config loading is in root because working directory of
		// the runtime is root
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

		authnUsecase := authn.NewUsecase(auth, dbQuerier, cacheStorer)
		foodUsecase := food.NewUsecase(dbQuerier, cacheStorer)
		usersUsecase := users.NewUsecase(dbQuerier)
		eventsUsecase := events.NewUsecase(dbQuerier, cacheStorer)

		grpcServer := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				// Should panic recoverer be put first or last?
				interceptors.RecoverPanic(),
				interceptors.ResponseError(),
				interceptors.ValidateToken(auth),
			),
		)
		apiauthv1.RegisterAuthServiceServer(grpcServer, handlers.NewAuthServer(auth, authnUsecase, usersUsecase))
		apifoodv1.RegisterFoodServiceServer(grpcServer, handlers.NewFoodServer(foodUsecase))

		worker := workers.New(authnUsecase, eventsUsecase)
		workerCtx := context.Background()
		workerTicker := time.NewTicker(10 * time.Second)
		defer workerTicker.Stop()

		go worker.Create(workerCtx, workerTicker.C)

		// This is not really idiomatic because
		// 	1.	Error should be returned instead of logged
		// 	2.	REST server can possibly start a split second before gRPC one does
		// 		and no mechanism to "delay" it
		go func() {
			if err := serveREST(cfg); err != nil {
				log.Fatal().Msg(err.Error())
			}
		}()

		port := fmt.Sprintf(":%v", cfg.GRPCServerPort)
		listener, _ := net.Listen("tcp", port)
		log.Info().Msgf("starting gRPC server at port %v", port)
		if err := grpcServer.Serve(listener); err != nil {
			return errors.New(errors.InternalServerError, "gRPC server failed to start")
		}

		return nil
	},
}

func serveREST(cfg *config.Config) error {
	grpcClient, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%v", cfg.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return errors.New(errors.InternalServerError,
			fmt.Sprintf("unable to create new gRPC client: %v", err))
	}

	gwMux := runtime.NewServeMux()
	if err = apiauthv1.RegisterAuthServiceHandler(context.Background(),
		gwMux, grpcClient); err != nil {
		return errors.New(errors.InternalServerError,
			fmt.Sprintf("unable to register auth handler to REST endpoint: %v", err))
	}
	if err = apifoodv1.RegisterFoodServiceHandler(context.Background(),
		gwMux, grpcClient); err != nil {
		return errors.New(errors.InternalServerError,
			fmt.Sprintf("unable to register food handler to REST endpoint: %v", err))
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.RESTServerPort),
		Handler: gwMux,
	}
	if err = gwServer.ListenAndServe(); err != nil {
		return errors.New(errors.InternalServerError, "REST server failed to start")
	}
	return nil
}
