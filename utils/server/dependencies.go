package server

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
)

type (
	controllersImpl struct {
	}
	servicesImpl struct {
	}
	repositoriesImpl struct {
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
	return controllersImpl{}
}

func resolveServices(repositories repositoriesImpl, clients clientsImpl) servicesImpl {
	return servicesImpl{}
}

func resolveRepositories(clients clientsImpl) repositoriesImpl {
	return repositoriesImpl{}
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
