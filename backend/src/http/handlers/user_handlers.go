package handlers

import (
	"log"
	"net/http"
)

func (app *Repository) Login(w http.ResponseWriter, r *http.Request) {
	// create request body struct
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	defer r.Body.Close()

	// decode request body into struct
	decodeRequestJSON(w, r, &body)

	log.Println(body)

	w.Write([]byte("Login"))
}
