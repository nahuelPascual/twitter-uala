package server

import (
	_ "github.com/lib/pq"

	"twitter-uala/src/controllers"
	"twitter-uala/utils/middlewares"
)

func (r routerImpl) registerRoutes() {
	healthCtrl := controllers.NewHealthController()
	r.router.GET("/health", healthCtrl.Health)

	clientsImpl := resolveClients()
	repositoriesImpl := resolveRepositories(clientsImpl)
	servicesImpl := resolveServices(repositoriesImpl, clientsImpl)
	controllersImpl := resolveControllers(servicesImpl, clientsImpl)

	api := r.router.Group("/api/v1", middlewares.CallerID)
	{
		tweets := api.Group("/tweets")
		{
			//tweets.GET("", nil) nice-to-have!
			tweets.POST("", nil)
		}

		users := api.Group("/users")
		{
			users.POST("/:id/follow", nil)
			// users.POST("/:id/unfollow", nil) nice-to-have!
			users.GET("/:id/timeline", nil)
		}

	}
}
