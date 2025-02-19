package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gatsu420/mary/app/api"
	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/config"
	"github.com/gatsu420/mary/app/handlers"
	"github.com/gatsu420/mary/app/interceptors"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/db/dbgen"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("failed to read config file: %v", err)
	}

	authSvc := auth.NewService(cfg.JWTSecret)
	authInterceptor, err := interceptors.NewAuthInterceptor(authSvc)
	if err != nil {
		log.Fatal().Msgf("failed to initialize auth interceptor: %v", err)
	}

	dbpool, _ := pgxpool.New(context.Background(), cfg.PostgresURL)
	defer dbpool.Close()

	dbQueries := dbgen.New(dbpool)
	foodUsecases := food.NewUsecases(dbQueries)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.ValidateTokenUnaryInterceptor),
	)
	api.RegisterAuthServiceServer(grpcServer, &handlers.AuthServer{
		Services: authSvc,
	})
	api.RegisterFoodServiceServer(grpcServer, &handlers.FoodServer{
		Usecases: foodUsecases,
	})

	port := fmt.Sprintf(":%v", cfg.GRPCServerPort)
	listener, _ := net.Listen("tcp", port)
	log.Info().Msgf("starting gRPC server at port %v", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("gRPC server failed to start")
	}
}
