package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Health(context *gin.Context)
}

type healthControllerImpl struct{}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}

func (healthCtrl *healthControllerImpl) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "health",
	})
}
