package main

import (
	"os"
	"net/http"
	"time"

	"know/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	apiRouter := router.NewRouter()
	apiRouter.AddRoutes()

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Minute,
		Addr:         GetPort(),
		Handler:      http.TimeoutHandler(apiRouter, 10*time.Minute, "SERVICE UNAVAILABLE"),
	}

	log.Println("Listening on :3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

//GetPort gets the port from heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
