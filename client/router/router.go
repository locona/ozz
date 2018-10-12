package router

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/api"
	"github.com/locona/ozz/client/api/server"
)

func HealthCheck(app *gin.Engine) {
	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
}

func V1(app *gin.Engine) {
	app.GET("/auth", api.GetAuth)
	app.POST("/auth", api.CreateAuth)

	{
		// server
		s := app.Group("/server")
		s.GET("/health", server.Health)
	}
}
