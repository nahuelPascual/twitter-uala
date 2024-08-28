package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"twitter-uala/utils/rest"
)

const UserIDKey = "user_id"

var CallerID = func(c *gin.Context) {
	userID := c.GetHeader("x-caller-id")
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: "x-caller-id is required"})
		return
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, rest.ErrorResponse{StatusCode: http.StatusBadRequest, ErrMsg: fmt.Sprintf("invalid x-caller-id: %s", err.Error())})
		return
	}

	c.Set(UserIDKey, id)
	c.Next()
}
