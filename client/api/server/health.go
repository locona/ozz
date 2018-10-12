package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	url := "http://localhost:4455/sample"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer X_ViMdTLAfhNrmUZlIehMQ0wtGUgp9Mjiho8mVABnxk.Ly7BbwC_vp1TSJ6wyvmE3wYutf1M8SWORKtj1kc5KnE")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	ctx.JSON(http.StatusOK, nil)
}
