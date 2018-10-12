package router

import "github.com/gin-gonic/gin"

func HealthCheck(app *gin.Engine) *gin.Engine {
	app.GET("/sample", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	return app
}
