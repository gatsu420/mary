package handlers

import (
	apiauthv1 "github.com/gatsu420/mary/api/gen/go/auth/v1"
	apifoodv1 "github.com/gatsu420/mary/api/gen/go/food/v1"
	"github.com/gatsu420/mary/app/auth"
	"github.com/gatsu420/mary/app/usecases/food"
	"github.com/gatsu420/mary/app/usecases/users"
)

type AuthServer struct {
	apiauthv1.UnimplementedAuthServiceServer
	auth         auth.Auth
	usersUsecase users.Usecase
}

func NewAuthServer(auth auth.Auth, usersUsecase users.Usecase) *AuthServer {
	return &AuthServer{
		auth:         auth,
		usersUsecase: usersUsecase,
	}
}

var _ apiauthv1.AuthServiceServer = (*AuthServer)(nil)

type FoodServer struct {
	apifoodv1.UnimplementedFoodServiceServer
	usecase food.Usecase
}

func NewFoodServer(usecase food.Usecase) *FoodServer {
	return &FoodServer{
		usecase: usecase,
	}
}

var _ apifoodv1.FoodServiceServer = (*FoodServer)(nil)
