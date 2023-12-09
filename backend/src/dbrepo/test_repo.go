package dbrepo

import (
	"errors"
	"shohinsan/MeetMinder/src/models"
)

type testDBRepo struct{}

func NewTestDBRepo() DatabaseRepository {
	return &testDBRepo{}
}

func (m *testDBRepo) CreateUser(user *models.User) (*models.User, error) {
	if user.ID == 0 {
		return nil, errors.New("invalid user id")
	}
	return user, nil
}

func (m *testDBRepo) GetUserByEmail(email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("invalid email")
	}

	return &models.User{}, nil
}

func (m *testDBRepo) GetUserById(id int64) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user id")
	}

	return &models.User{}, nil
}
