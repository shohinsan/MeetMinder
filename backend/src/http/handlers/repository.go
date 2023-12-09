package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shohinsan/MeetMinder/src/dbrepo"
)

type Repository struct {
	DB dbrepo.DatabaseRepository
}

func NewRepository(db dbrepo.DatabaseRepository) *Repository {
	return &Repository{
		DB: db,
	}
}

var Repo *Repository

func NewHandlers(repo *Repository) {
	Repo = repo
}

func decodeRequestJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}
