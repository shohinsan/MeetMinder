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
	if user.Email == "error@test.loc" {
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

// CreateMeeting implements DatabaseRepository.
func (*testDBRepo) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	panic("unimplemented")
}

// DeleteMeeting implements DatabaseRepository.
func (*testDBRepo) DeleteMeeting(id int64) error {
	panic("unimplemented")
}

// GetMeetingById implements DatabaseRepository.
func (*testDBRepo) GetMeetingById(id int64) (*models.Meeting, error) {
	panic("unimplemented")
}

// GetMeetingsForUser implements DatabaseRepository.
func (*testDBRepo) GetMeetingsForUser(userId int64) ([]*models.Meeting, error) {
	panic("unimplemented")
}

// UpdateMeeting implements DatabaseRepository.
func (*testDBRepo) UpdateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	panic("unimplemented")
}
