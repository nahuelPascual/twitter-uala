package server

import (
	_ "github.com/lib/pq"

	"twitter-uala/src/controllers"
)

func (r routerImpl) registerRoutes() {
	healthCtrl := controllers.NewHealthController()
	r.router.GET("/health", healthCtrl.Health)

	clientsImpl := resolveClients()
	repositoriesImpl := resolveRepositories(clientsImpl)
	servicesImpl := resolveServices(repositoriesImpl, clientsImpl)
	controllersImpl := resolveControllers(servicesImpl, clientsImpl)

	api := r.router.Group("/api/v1")
	{
		tweets := api.Group("/tweets")
		{
			//tweets.GET("", nil) nice-to-have!
			tweets.POST("", controllersImpl.TweetsController.Publish)
		}

		users := api.Group("/users")
		{
			users.POST("/:username/follow", controllersImpl.UsersController.Follow)
			// users.POST("/:username/unfollow", controllersImpl.UsersController.Unfollow) nice-to-have!
		}

		api.GET("/timeline", controllersImpl.TimelineController.GetTimeline)

	}
}
