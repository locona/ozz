package router

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/api/hydra"
	"github.com/locona/ozz/client/api/server"
)

func V1(app *gin.Engine) {
	app.GET("/auth", hydra.GetAuth)
	app.POST("/auth", hydra.CreateAuth)

	{
		h := app.Group("/health")
		h.GET("/alive", server.Alive)
		h.GET("/ready", server.Ready)
	}

	{
		sess := app.Group("/session")
		sess.POST("/signup", server.Signup)
	}
}
