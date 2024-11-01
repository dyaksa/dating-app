//go:build wireinject
// +build wireinject

package server

import (
	"github.com/dyaksa/dating-app/internal/app/http/middleware"
	"github.com/dyaksa/dating-app/internal/app/http/route"
	"github.com/dyaksa/dating-app/internal/config"
	authDelivery "github.com/dyaksa/dating-app/internal/modules/auth/v1/delivery"
	authRepo "github.com/dyaksa/dating-app/internal/modules/auth/v1/repository"
	authUsecase "github.com/dyaksa/dating-app/internal/modules/auth/v1/usecase"
	purchasesDelivery "github.com/dyaksa/dating-app/internal/modules/purchases/v1/delivery"
	purchasesRepo "github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository"
	purchasesUsecase "github.com/dyaksa/dating-app/internal/modules/purchases/v1/usecase"
	swipesDelivery "github.com/dyaksa/dating-app/internal/modules/swipes/v1/delivery"
	swipesRepo "github.com/dyaksa/dating-app/internal/modules/swipes/v1/repository"
	swipesUsecase "github.com/dyaksa/dating-app/internal/modules/swipes/v1/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.NewPostgres,
)

var repositorySet = wire.NewSet(
	authRepo.NewAuthRepository,
	swipesRepo.NewSwipesRepository,
	purchasesRepo.NewPurchasesRepository,
)

var usecaseSet = wire.NewSet(
	authUsecase.NewAuthUsecaseImpl,
	swipesUsecase.NewSwipeUsecase,
	purchasesUsecase.NewPurchasesUsecase,
)

var handlerSet = wire.NewSet(
	authDelivery.NewRestAuthHandler,
	swipesDelivery.NewRestSwipesHandler,
	purchasesDelivery.NewRestPurchasesHandler,
)

var authMiddleware = wire.NewSet(
	middleware.NewAuthMiddleware,
)

func InitializeServer() *gin.Engine {
	wire.Build(configSet, repositorySet, usecaseSet, handlerSet, authMiddleware, route.NewRoute, config.NewApp)
	return nil
}
