package main

import (
	"context"
	"fmt"
	"net"
	"time"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/cache"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/events"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/app/usecases/users"
	"github.com/gatsu420/mary/app/workers"
	"github.com/gatsu420/mary/common/config"
	"github.com/gatsu420/mary/dependency/pgdep"
	"github.com/gatsu420/mary/dependency/valkeydep"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		log.Fatal().Msgf("failed to read config file: %v", err)
	}

	dbPool, err := pgdep.NewPool(cfg.PostgresDSN)
	if err != nil {
		log.Fatal().Msgf("failed to create DB connection: %v", err)
	}
	defer dbPool.Close()
	dbQuerier := repository.New(dbPool)

	valkeyClient, err := valkeydep.New(cfg.CacheAddress)
	if err != nil {
		log.Fatal().Msgf("failed to create cache connection: %v", err)
	}
	defer valkeyClient.Close()
	cacheStorer := cache.New(valkeyClient)

	auth := auth.NewAuth(cfg.JWTSecret)

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
	apiauthv1.RegisterAuthServiceServer(grpcServer, handlers.NewAuthServer(auth, usersUsecase))
	apifoodv1.RegisterFoodServiceServer(grpcServer, handlers.NewFoodServer(foodUsecase))

	worker := workers.New(eventsUsecase)
	workerCtx := context.Background()
	workerTicker := time.NewTicker(10 * time.Second)
	defer workerTicker.Stop()
	go worker.Create(workerCtx, workerTicker.C)

	port := fmt.Sprintf(":%v", cfg.GRPCServerPort)
	listener, _ := net.Listen("tcp", port)
	log.Info().Msgf("starting gRPC server at port %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("gRPC server failed to start")
	}
}
