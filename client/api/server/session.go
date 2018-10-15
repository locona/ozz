package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func Signup(ctx *gin.Context) {
	url := "http://localhost:24455/session/signup"
	resp, err := post(url, "")
	pp.Println(err)
	if resp.StatusCode >= 400 || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
