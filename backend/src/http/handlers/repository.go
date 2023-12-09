package handlers

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

var Repo *Repository

func NewHandlers(repo *Repository) {
	Repo = repo
}
