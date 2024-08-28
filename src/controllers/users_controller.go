package controllers

import (
	"fmt"
	"log"
	"net/http"

	"twitter-uala/src/controllers/dto"
	"twitter-uala/src/errors"
	"twitter-uala/src/errors/validations"
	"twitter-uala/src/services"
	"twitter-uala/src/utils/rest"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService services.UsersService
}

func (c *UsersController) Follow(ctx *gin.Context) {
	var headers dto.UserIDHeader
	if err := ctx.ShouldBindHeader(&headers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: err.Error()})
		return
	}

	var params dto.FollowUserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: err.Error()})
		return
	}

	if err := c.UsersService.Follow(headers.UserID, params.Username); err != nil {
		log.Println(fmt.Sprintf("error following user %s: %s", params.Username, err.Error()))
		switch err.(type) {
		case errors.EntityNotFoundError:
			ctx.AbortWithStatusJSON(http.StatusNotFound, rest.ErrorResponse{StatusCode: http.StatusNotFound, ErrMsg: err.Error()})
			return
		case validations.FollowingHimSelfError:
			ctx.AbortWithStatusJSON(http.StatusConflict, rest.ErrorResponse{StatusCode: http.StatusConflict, ErrMsg: err.Error()})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, rest.ErrorResponse{StatusCode: http.StatusInternalServerError, ErrMsg: err.Error()})
			return
		}
	}

	ctx.Status(http.StatusOK)
}
