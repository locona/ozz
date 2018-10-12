package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/locona/ozz/client/pkg/hydra"
)

func GetAuth(ctx *gin.Context) {
	cli, err := hydra.New()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := cli.Token()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errors.New("invalid")})
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func CreateAuth(ctx *gin.Context) {
	cli, err := hydra.New()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = cli.CreateClient()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}
