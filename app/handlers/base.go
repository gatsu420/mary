package handlers

import (
	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/app/usecases/users"
	"github.com/gatsu420/mary/auth"
)

type AuthServer struct {
	apiauthv1.UnimplementedAuthServiceServer
	Auth         auth.Auth
	UsersUsecase users.Usecase
}

func NewAuthServer(auth auth.Auth, usersUsecase users.Usecase) *AuthServer {
	return &AuthServer{
		Auth:         auth,
		UsersUsecase: usersUsecase,
	}
}

type FoodServer struct {
	apifoodv1.UnimplementedFoodServiceServer
	Usecase food.Usecase
}

func NewFoodServer(usecase food.Usecase) *FoodServer {
	return &FoodServer{
		Usecase: usecase,
	}
}
