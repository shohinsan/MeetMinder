package dbrepo

import "shohinsan/MeetMinder/src/models"

type DatabaseRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int64) (*models.User, error)
}
