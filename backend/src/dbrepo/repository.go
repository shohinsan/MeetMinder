package dbrepo

import "shohinsan/MeetMinder/src/models"

type DatabaseRepository interface {
	// Create a new user, returns the new user or an error
	CreateUser(user *models.User) (*models.User, error)

	// Get a user by email, returns the user or an error
	GetUserByEmail(email string) (*models.User, error)

	// Get a user by id, returns the user or an error
	GetUserById(id int64) (*models.User, error)
}
