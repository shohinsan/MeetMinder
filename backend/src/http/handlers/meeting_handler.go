package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shohinsan/MeetMinder/src/models"
	"time"

	"github.com/gorilla/mux"
)

func (app *Repository) CreateMeeting(w http.ResponseWriter, r *http.Request) {
	// create request body struct
	body := struct {
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		StartTime    time.Time `json:"start_time"`
		EndTime      time.Time `json:"end_time"`
		Participants []string  `json:"participants"`
	}{}

	// decode request body into struct
	decodeRequestJSON(w, r, &body)

	// validate request body
	if body.Title == "" || body.StartTime.IsZero() || body.EndTime.IsZero() {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// get host id
	hostUsername := mux.Vars(r)["username"]
	host, err := app.DB.GetUserByUsername(hostUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// get authenticated user's username
	claims := r.Context().Value(ContextKey{}).(map[string]interface{})
	creatorUsername := claims["username"].(string)

	// get user from database
	user, err := app.DB.GetUserByUsername(creatorUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// Create a list of participants
	participants := make(map[string]bool)
	participants[host.Username] = true
	participants[user.Username] = true

	for _, username := range body.Participants {
		participants[username] = true
	}

	// TODO: Save participants to database

	// create meeting
	meeting := &models.Meeting{
		Title:       body.Title,
		Description: body.Description,
		HostID:      host.ID,
		CreatorID:   user.ID,
		StartTime:   body.StartTime,
		EndTime:     body.EndTime,
	}

	// save meeting to database
	meeting, err = app.DB.CreateMeeting(meeting)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// send meeting to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meeting)
}
