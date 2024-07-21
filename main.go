package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"skeleton-web-app/src/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// Pass the router to the middleware.Router function
	httpMiddleware := middleware.Router(router)

	// Check environment variables for TLS mode
	tlsMode := os.Getenv("TLS_MODE") == "true"
	httpVersion := os.Getenv("HTTP_VERSION") == "2"

	if tlsMode || httpVersion {
		log.Println("Starting server in TLS mode...")
		if err := httpMiddleware.RunTLS(":8080", "./certs/server.pem", "./certs/server.key"); err != nil {
			log.Fatal(fmt.Errorf("error starting TLS server: %v", err))
		}
		return
	}

	log.Println("Starting server in non-TLS mode...")
	// Create an HTTP server instance
	srv := &http.Server{
		Addr:    ":8080",
		Handler: httpMiddleware,

		// Set timeout to mitigate CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 3 * time.Second,
	}

	// Start the HTTP server
	log.Fatal(srv.ListenAndServe())
}
