package route

import (
	"github.com/dyaksa/dating-app/internal/app/http/middleware"
	authHandler "github.com/dyaksa/dating-app/internal/modules/auth/v1/delivery"
	purchasesHandler "github.com/dyaksa/dating-app/internal/modules/purchases/v1/delivery"
	swipesHandler "github.com/dyaksa/dating-app/internal/modules/swipes/v1/delivery"
	"github.com/gin-gonic/gin"
)

type ConfigRoute struct {
	authMiddleware   *middleware.AuthMiddleware
	authHandler      *authHandler.RestAuthHandler
	swipesHandler    *swipesHandler.RestSwipesHandler
	purchasesHandler *purchasesHandler.RestPurchasesHandler
}

func NewRoute(
	authMiddleware *middleware.AuthMiddleware,
	authHandler *authHandler.RestAuthHandler,
	swipesHandler *swipesHandler.RestSwipesHandler,
	purchasesHandler *purchasesHandler.RestPurchasesHandler,

) *ConfigRoute {
	return &ConfigRoute{
		authMiddleware:   authMiddleware,
		authHandler:      authHandler,
		swipesHandler:    swipesHandler,
		purchasesHandler: purchasesHandler,
	}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	c.healthCheck(app)
	c.authRouteApiV1(app)
	c.swipesRouteApiV1(app)
	c.purchasesRouteApiV1(app)
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

func (c *ConfigRoute) swipesRouteApiV1(app *gin.Engine) {
	v1 := app.Group("/api/v1", c.authMiddleware.Authenticated())
	{
		v1.GET("/user/swipes", c.swipesHandler.GetSwipeUser)
		v1.POST("/user/swipes", c.swipesHandler.SwipeUser)
	}
}

func (c *ConfigRoute) purchasesRouteApiV1(app *gin.Engine) {
	v1 := app.Group("/api/v1", c.authMiddleware.Authenticated())
	{
		v1.POST("/user/purchases", c.purchasesHandler.Purchases)
	}
}
