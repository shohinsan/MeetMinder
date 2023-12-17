package dbrepo

import "shohinsan/MeetMinder/src/models"

type DatabaseRepository interface {
	// Create a new user, returns the new user or an error
	CreateUser(user *models.User) (*models.User, error)

	// Get a user by email, returns the user or an error
	GetUserByEmail(email string) (*models.User, error)

	// Get a user by id, returns the user or an error
	GetUserById(id int64) (*models.User, error)

	// Get a user by username, returns the user or an error
	GetUserByUsername(username string) (*models.User, error)

	// Create a new meeting, returns the new meeting or an error
	CreateMeeting(meeting *models.Meeting) (*models.Meeting, error)

	// Get a meeting by id, returns the meeting or an error
	GetMeetingById(id int64) (*models.Meeting, error)

	// Get all meetings for a user, returns a list of meetings or an error
	GetMeetingsForUser(userId int64) ([]*models.Meeting, error)

	// Update a meeting, returns the updated meeting or an error
	UpdateMeeting(meeting *models.Meeting) (*models.Meeting, error)

	// Delete a meeting by id, returns an error if something goes wrong
	DeleteMeeting(id int64) error
}
