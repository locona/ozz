package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/infra"
	"github.com/locona/ozz/client/pkg/hydra"
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

func get(url, token string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	bearer := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", bearer)
	client := new(http.Client)
	return client.Do(req)
}

func token() (string, error) {
	cache := hydra.NewCache(infra.Redis)
	return cache.Get()
}
