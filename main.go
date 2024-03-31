package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	GtmID = os.Getenv("MYSTIC_CASE_GTM_ID")
)

func main() {
	log.Print("starting server...")

	initRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
