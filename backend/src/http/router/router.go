package router

import (
	"shohinsan/MeetMinder/src/http/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Set up routes
	router.HandleFunc("/api/login", handlers.Repo.Login).Methods("POST")

	return router
}
