package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetsController struct{}

func (c *TweetsController) Publish(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
