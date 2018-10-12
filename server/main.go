package main

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/server/router"
)

func main() {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	router.HealthCheck(api)

	api.Run(":9999")
}
