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

	// validate request body
	if body.Email == "" || body.Password == "" {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	// get user from database
	user, err := app.DB.GetUserByEmail(body.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// compare password with hash
	err = app.hashRepo.Compare(user.Password, body.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	// Create auth token
	payload := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	}

	token, err := app.authTokenRepo.CreateToken(payload)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// send token to client
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}
