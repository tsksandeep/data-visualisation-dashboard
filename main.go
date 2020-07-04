package main

import (
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
		Addr:         ":3000",
		Handler:      http.TimeoutHandler(apiRouter, 10*time.Minute, "SERVICE UNAVAILABLE"),
	}

	log.Println("Listening on :3000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
