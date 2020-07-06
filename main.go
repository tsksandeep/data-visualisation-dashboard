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

	log "github.com/ctrlrsf/logdna"
)

//CreateStores creates all the stores
func createStores(postgresDB *db.DB, logDNAClient *log.Client) (*models.Stores, error) {

	var stores models.Stores

	accountStore, err := postgres.NewAccountStore(postgresDB, logDNAClient)
	if err != nil {
		return nil, err
	}

	stores.AccountStore = accountStore

	return &stores, nil

}


func main() {

	logDNAConfig := log.Config{
		APIKey:     os.Getenv("LOGDNA_KEY"),
		LogFile:    "know.log",
		FlushLimit: 1,
	}

	logDNAClient := log.NewClient(logDNAConfig)
	defer logDNAClient.Close()

	postgresDB, err := db.NewDB(logDNAClient)
	if err != nil {
		logDNAClient.Log(time.Now(), err.Error())
		os.Exit(1)
	}

	defer postgresDB.Close()

	stores, err := createStores(postgresDB, logDNAClient)
	if err != nil {
		logDNAClient.Log(time.Now(), fmt.Sprintf("creating stores failed: %s", err.Error()))
		os.Exit(1)
	}

	apiRouter := router.NewRouter()
	apiRouter.AddRoutes(stores)
	port := os.Getenv("PORT")
	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Minute,
		Addr:         port,
		Handler:      http.TimeoutHandler(apiRouter, 10*time.Minute, "SERVICE UNAVAILABLE"),
	}

	logDNAClient.Log(time.Now(), fmt.Sprintf("Listening on %s", port))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logDNAClient.Log(time.Now(), err.Error())
		os.Exit(1)
	}
}
