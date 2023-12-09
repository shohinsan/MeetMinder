package handlers

import "shohinsan/MeetMinder/src/dbrepo"

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
