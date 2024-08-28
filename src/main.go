package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"twitter-uala/src/server"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func getHostPort() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	log.Println("Starting twitter-uala")
	gin.SetMode(gin.ReleaseMode)

	if err := runMigrations(); err != nil {
		log.Fatalf(err.Error())
	}

	router := server.NewRouter(getHostPort())
	router.StartUp()
}

func runMigrations() error {
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	connString := fmt.Sprintf("user=%s password=%s database=%s host=%s sslmode=disable", username, password, dbName, host)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("error opening sql connection. %s", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating postgres instance. %s", err.Error())
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dbName,
		driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migrations with instance. %s", err.Error())
	}

	if err = migrations.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) && !errors.Is(err, migrate.ErrNilVersion) {
			return fmt.Errorf("error running migrations. %s", err.Error())
		}
	}

	v, dirty, err := migrations.Version()
	if err != nil {
		return fmt.Errorf("error getting db version: %s", err.Error())
	}
	log.Printf("database version: (%d, %t)\n", v, dirty)

	return nil
}
