package main

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/config"
	"github.com/locona/ozz/client/infra"
	"github.com/locona/ozz/client/router"
)

func init() {
	config.Init()
	infra.InitRedis()
}

func main() {
	api := gin.New()
	api.HandleMethodNotAllowed = true
	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	router.V1(api)

	api.Run(":9998")
}
