package route

import "github.com/gin-gonic/gin"

type ConfigRoute struct{}

func NewRoute() *ConfigRoute {
	return &ConfigRoute{}
}

func (c *ConfigRoute) Setup(app *gin.Engine) {
	c.healthCheck(app)
}

func (c *ConfigRoute) healthCheck(app *gin.Engine) {
	app.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK!",
		})
	})
}
