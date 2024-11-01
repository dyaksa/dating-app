//go:build wireinject
// +build wireinject

package server

import (
	"github.com/dyaksa/dating-app/internal/app/http/route"
	"github.com/dyaksa/dating-app/internal/config"
	authDelivery "github.com/dyaksa/dating-app/internal/modules/auth/v1/delivery"
	authRepo "github.com/dyaksa/dating-app/internal/modules/auth/v1/repository"
	authUsecase "github.com/dyaksa/dating-app/internal/modules/auth/v1/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.NewPostgres,
)

var repositorySet = wire.NewSet(
	authRepo.NewAuthRepository,
)

var usecaseSet = wire.NewSet(
	authUsecase.NewAuthUsecaseImpl,
)

var handlerSet = wire.NewSet(
	authDelivery.NewRestAuthHandler,
)

func InitializeServer() *gin.Engine {
	wire.Build(configSet, repositorySet, usecaseSet, handlerSet, route.NewRoute, config.NewApp)
	return nil
}
