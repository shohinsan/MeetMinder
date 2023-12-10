package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shohinsan/MeetMinder/src/models"
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

	// TODO: Implement input validation
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

	token, err := app.AuthTokenRepo.CreateToken(payload)
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

func (app *Repository) Register(w http.ResponseWriter, r *http.Request) {
	// create request body struct
	body := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	defer r.Body.Close()

	// decode request body into struct
	decodeRequestJSON(w, r, &body)

	// TODO: Implement input validation
	// validate request body
	if body.Email == "" || body.Username == "" || body.Password == "" {
		http.Error(w, "Invalid email, username, or password", http.StatusBadRequest)
		return
	}

	// hash password
	hash, err := app.hashRepo.Hash(body.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// create user
	user := &models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: hash,
	}

	// save user to database
	user, err = app.DB.CreateUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// Create auth token
	payload := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	}

	token, err := app.AuthTokenRepo.CreateToken(payload)
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

func (app *Repository) User(w http.ResponseWriter, r *http.Request) {
	// get user from context
	claims := r.Context().Value(ContextKey{}).(map[string]interface{})

	// send user to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}
