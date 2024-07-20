package main

import (
	"fmt"
	"log"
	"net/http"
	"skeleton-web-app/src/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	initRouter := middleware.Router(router)

	// Attempt to run the server
	if err := initRouter.Run(":8080"); err != nil {
		log.Fatal(fmt.Errorf("error starting server: %v", err))
	}

	// Create an HTTP server instance
	srv := &http.Server{
		Addr:    ":8080",
		Handler: initRouter,
		// Set timeout to mitigate CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Start the HTTP server
	log.Fatal(srv.ListenAndServe())
}
