package main

import (
	"log"
	"net/http"
	"shohinsan/MeetMinder/src/http/handlers"
	"shohinsan/MeetMinder/src/http/router"

	"github.com/joho/godotenv"
)

const PORT = ":8080"

func main() {
	godotenv.Load()

	// Set up HTTP handlers
	handlers.NewHandlers(
		handlers.NewRepository(),
	)

	// Set up HTTP router
	r := router.NewRouter()

	// Set up HTTP server
	server := &http.Server{
		Addr:    PORT,
		Handler: r,
	}

	log.Printf("Starting server on port %s", PORT)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
