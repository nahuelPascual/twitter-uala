package server

import (
	"fmt"

	"twitter-uala/src/utils/middlewares"

	"github.com/gin-gonic/gin"
)

type routerImpl struct {
	router *gin.Engine
	port   string
}

type Router interface {
	StartUp()
	GetRouter() *gin.Engine
}

func NewRouter(port string) *routerImpl {
	return &routerImpl{
		router: gin.Default(),
		port:   port,
	}
}

func (r routerImpl) StartUp() {
	r.router.Use(
		gin.Recovery(),
		gin.Logger(),
		middlewares.CORSMiddleware(),
	)
	r.registerRoutes()
	if err := r.router.Run(r.port); err != nil {
		fmt.Errorf("unable to start router error: %!v(MISSING)", err)
		panic(err)
	}
}

func (r routerImpl) GetRouter() *gin.Engine {
	return r.router
}
