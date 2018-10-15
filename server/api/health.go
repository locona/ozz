package api

import "github.com/gin-gonic/gin"

func Alive(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func Ready(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
