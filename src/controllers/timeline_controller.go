package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"twitter-uala/src/controllers/dto"
	"twitter-uala/src/services"
	"twitter-uala/src/utils/rest"
)

type TimelineController struct {
	TimelineService services.TimelineService
}

func (c TimelineController) GetTimeline(ctx *gin.Context) {
	var headers dto.UserIDHeader
	if err := ctx.ShouldBindHeader(&headers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: err.Error()})
		return
	}

	tweets, err := c.TimelineService.ResolveTimeline(headers.UserID)
	if err != nil {
		log.Println(fmt.Sprintf("error resolving timeline for user %d: %s", headers.UserID, err.Error()))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, rest.ErrorResponse{StatusCode: http.StatusInternalServerError, ErrMsg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.TimelineResponse{Tweets: tweets, TweetsCount: len(tweets)})
}
