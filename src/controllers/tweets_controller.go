package controllers

import (
	"fmt"
	"log"
	"net/http"

	"twitter-uala/src/controllers/dto"
	"twitter-uala/src/services"
	"twitter-uala/src/utils/rest"

	"github.com/gin-gonic/gin"
)

type TweetsController struct {
	TweetsService services.TweetsService
}

func (c *TweetsController) Publish(ctx *gin.Context) {
	var headers dto.UserIDHeader
	if err := ctx.ShouldBindHeader(&headers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: err.Error()})
		return
	}

	var request dto.Tweet
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: err.Error()})
		return
	}

	if err := c.TweetsService.Publish(ctx, request.ToDomain(headers.UserID)); err != nil {
		log.Println(fmt.Sprintf("error creating tweet: %s", err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, rest.ErrorResponse{StatusCode: http.StatusInternalServerError, ErrMsg: err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
