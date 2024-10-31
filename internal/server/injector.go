//go:build wireinject
// +build wireinject

package server

import (
	"github.com/dyaksa/dating-app/internal/app/http/route"
	"github.com/dyaksa/dating-app/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.NewPostgres,
)

func InitializeServer() *gin.Engine {
	wire.Build(route.NewRoute, config.NewApp)
	return nil
}
