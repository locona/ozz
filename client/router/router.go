package router

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/api"
	"github.com/locona/ozz/client/api/server"
)

func V1(app *gin.Engine) {
	app.GET("/auth", api.GetAuth)
	app.POST("/auth", api.CreateAuth)

	{
		// server
		h := app.Group("/health")
		h.GET("/alive", server.Alive)
		h.GET("/ready", server.Ready)
	}
}
