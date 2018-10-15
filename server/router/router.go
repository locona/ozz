package router

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/server/api"
)

func V1(app *gin.Engine) {
	{
		h := app.Group("/health")
		h.GET("/alive", api.Alive)
		h.GET("/ready", api.Ready)
	}

	{
		sess := app.Group("/session")
		sess.POST("/signup")
	}
}
