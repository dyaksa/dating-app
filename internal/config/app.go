package config

import (
	"github.com/dyaksa/dating-app/internal/app/http/route"
	"github.com/gin-gonic/gin"
)

func NewApp(route *route.ConfigRoute) *gin.Engine {
	var app = gin.Default()

	route.Setup(app)

	return app
}
