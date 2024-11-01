package route

import (
	authHandler "github.com/dyaksa/dating-app/internal/modules/auth/v1/delivery"
	"github.com/gin-gonic/gin"
)

type ConfigRoute struct {
	authHandler *authHandler.RestAuthHandler
}

func NewRoute(
	authHandler *authHandler.RestAuthHandler,
) *ConfigRoute {
	return &ConfigRoute{
		authHandler: authHandler,
	}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	c.healthCheck(app)
	c.authRouteApiV1(app)
}

func (c *ConfigRoute) healthCheck(app *gin.Engine) {
	app.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK!",
		})
	})
}

func (c *ConfigRoute) authRouteApiV1(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	{
		v1.POST("/auth/register", c.authHandler.Register)
		v1.POST("/auth/login", c.authHandler.Login)
	}
}
