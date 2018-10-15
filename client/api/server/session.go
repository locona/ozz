package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	url := "http://localhost:4455/session/signup"
	resp, err := post(url, "")
	defer resp.Body.Close()
	if resp.StatusCode >= 400 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
