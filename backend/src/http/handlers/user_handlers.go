package handlers

import "net/http"

func (app *Repository) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}
