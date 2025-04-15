package main

import (
	"context"
	"fmt"
	"net"

	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/app/repository"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/app/usecases/users"
	"github.com/gatsu420/mary/auth"
	"github.com/gatsu420/mary/common/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		log.Fatal().Msgf("failed to read config file: %v", err)
	}

	dbpool, _ := pgxpool.New(context.Background(), cfg.PostgresURL)
	defer dbpool.Close()
	dbQueries := repository.New(dbpool)

	auth := auth.NewAuth(cfg.JWTSecret)

	foodUsecase := food.NewUsecase(dbQueries)
	usersUsecase := users.NewUsecase(dbQueries)

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

	port := fmt.Sprintf(":%v", cfg.GRPCServerPort)
	listener, _ := net.Listen("tcp", port)
	log.Info().Msgf("starting gRPC server at port %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("gRPC server failed to start")
	}
}
