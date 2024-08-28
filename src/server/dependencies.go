package server

import (
	"fmt"
	"net/url"
	"os"

	"twitter-uala/src/controllers"
	"twitter-uala/src/repositories"
	"twitter-uala/src/services"

	"github.com/jmoiron/sqlx"
)

type (
	controllersImpl struct {
		TweetsController controllers.TweetsController
	}
	servicesImpl struct {
		TweetsService services.TweetsService
	}
	repositoriesImpl struct {
		TweetsRepository repositories.TweetsRepository
	}
	clientsImpl struct {
		dbClient *sqlx.DB
	}
)

func resolveClients() clientsImpl {
	return clientsImpl{
		dbClient: resolveSQLClient(),
	}
}

func resolveControllers(services servicesImpl, clients clientsImpl) controllersImpl {
	return controllersImpl{
		TweetsController: controllers.TweetsController{TweetsService: services.TweetsService},
	}
}

func resolveServices(repositories repositoriesImpl, clients clientsImpl) servicesImpl {
	return servicesImpl{
		TweetsService: services.NewTweetsService(repositories.TweetsRepository),
	}
}

func resolveRepositories(clients clientsImpl) repositoriesImpl {
	return repositoriesImpl{
		TweetsRepository: repositories.NewTweetsRepository(clients.dbClient),
	}
}

func resolveSQLClient() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	var connString string
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		connString = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, url.QueryEscape(pass), host, dbName)
	} else {
		connString = fmt.Sprintf("user=%s password=%s database=%s host=%s sslmode=disable", user, pass, dbName, host)
	}

	client, err := sqlx.Open("postgres", connString)
	if err != nil {
		panic(fmt.Sprintf("error to create db client. %s", err.Error()))
	}

	return client
}
