package main

import (
	"helios-proxy/pkg/auth"
	"helios-proxy/pkg/proxy"
	"log"
	"net/http"
)

func main() {
	// Initialize the proxy server
	router := proxy.NewRouter()

	// Set up middleware for authorization
	router.Use(auth.JWTMiddleware)
	router.Use(auth.OAuthMiddleware)

	// Start the server
	log.Println("Starting Helios Proxy on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
