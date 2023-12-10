package main

import (
	"log"
	"net/http"
	"shohinsan/MeetMinder/src/dbrepo"
	"shohinsan/MeetMinder/src/http/handlers"
	"shohinsan/MeetMinder/src/http/router"
	"shohinsan/MeetMinder/src/services/hashrepo"
	"shohinsan/MeetMinder/src/services/tokenrepo"

	"github.com/joho/godotenv"
)

const PORT = ":8080"

func main() {
	godotenv.Load()

	// Set up Database Repository
	db := dbrepo.NewTestDBRepo()

	// Set up Hash Repository
	hr := hashrepo.NewTestHashRepo()

	// Set up Token Repository
	tr := tokenrepo.NewTestTokenRepo()

	// Set up HTTP handlers
	handlers.NewHandlers(
		handlers.NewRepository(db, hr, tr),
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
