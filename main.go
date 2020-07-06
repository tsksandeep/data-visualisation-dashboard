package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"know/db"
	"know/models"
	"know/models/postgres"
	"know/router"

	"github.com/logdna/logdna-go/logger"
)

//CreateStores creates all the stores
func createStores(postgresDB *db.DB, log *logger.Logger) (*models.Stores, error) {
	var stores models.Stores
	accountStore, err := postgres.NewAccountStore(postgresDB, log)
	if err != nil {
		return nil, err
	}

	stores.AccountStore = accountStore
	return &stores, nil
}

func main() {

	log := logger.CreateLogger(logger.Options{
		Level:         "fatal",
		Hostname:      "Know",
		App:           "know-dash",
		Env:           "production",
		FlushInterval: 1 * time.Second,
	}, os.Getenv("LOGDNA_KEY"))

	defer log.Close()

	postgresDB, err := db.NewDB(log)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer postgresDB.Close()

	stores, err := createStores(postgresDB, log)
	if err != nil {
		log.Fatal(fmt.Sprintf("creating stores failed: %s", err.Error()))
	}

	apiRouter := router.NewRouter()
	apiRouter.AddRoutes(stores, log)
	port := ":" + os.Getenv("PORT")
	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Minute,
		Addr:         port,
		Handler:      http.TimeoutHandler(apiRouter, 10*time.Minute, "SERVICE UNAVAILABLE"),
	}

	log.Info(fmt.Sprintf("Listening on %s", port))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
