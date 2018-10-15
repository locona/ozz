package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alive(ctx *gin.Context) {
	token, err := token()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://localhost:4455/health/alive"
	resp, err := get(url, token)
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func Ready(ctx *gin.Context) {
	token, err := token()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	url := "http://localhost:4455/health/ready"
	resp, err := get(url, token)
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
