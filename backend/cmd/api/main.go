package main

import (
	"log"
	"net/http"
	"shohinsan/MeetMinder/src/dbrepo"
	"shohinsan/MeetMinder/src/http/handlers"
	"shohinsan/MeetMinder/src/http/router"

	"github.com/joho/godotenv"
)

const PORT = ":8080"

func main() {
	godotenv.Load()

	// Set up Database Repository
	db := dbrepo.NewTestDBRepo()

	// Set up HTTP handlers
	handlers.NewHandlers(
		handlers.NewRepository(db),
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
