package main

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/router"
)

func main() {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	router.HealthCheck(api)
	router.V1(api)

	api.Run(":9998")
}
