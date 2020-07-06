package main

import (
	"net/http"
	"os"
	"time"

	"know/db"
	"know/models"
	"know/models/postgres"
	"know/router"

	log "github.com/sirupsen/logrus"
)

//CreateStores creates all the stores
func createStores(postgresDB *db.DB) (*models.Stores, error) {

	var stores models.Stores

	accountStore, err := postgres.NewAccountStore(postgresDB)
	if err != nil {
		return nil, err
	}

	stores.AccountStore = accountStore

	return &stores, nil

}

//GetPort gets the port from heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func main() {

	postgresDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer postgresDB.Close()

	stores, err := createStores(postgresDB)
	if err != nil {
		log.Fatal("creating stores failed: ", err)
	}

	apiRouter := router.NewRouter()
	apiRouter.AddRoutes(stores)

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Minute,
		Addr:         getPort(),
		Handler:      http.TimeoutHandler(apiRouter, 10*time.Minute, "SERVICE UNAVAILABLE"),
	}

	log.Println("Listening on :3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
