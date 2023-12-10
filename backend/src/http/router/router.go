package router

import (
	"shohinsan/MeetMinder/src/http/handlers"
	"shohinsan/MeetMinder/src/http/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Set up middleware
	mr := middleware.NewMiddlewareRepository(handlers.Repo)

	// Set up routes
	r.HandleFunc("/api/login", handlers.Repo.Login).Methods("POST")
	r.HandleFunc("/api/register", handlers.Repo.Register).Methods("POST")

	// Set up protected routes
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/user", handlers.Repo.User).Methods("GET")
	s.Use(mr.RequiresAuthentication)

	return r
}
