package dbrepo

import (
	"context"
	"database/sql"
	"shohinsan/MeetMinder/src/models"
	"time"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) DatabaseRepository {
	return &postgresRepo{
		db: db,
	}
}

// CreateUser implements DatabaseRepository.
func (m *postgresRepo) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (email, username, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newID int64

	err := m.db.QueryRowContext(
		ctx, query, user.Email, user.Username, user.FirstName, user.LastName, user.Password,
	).Scan(&newID)
	if err != nil {
		return nil, err
	}

	user.ID = newID

	return user, nil
}

// GetUserByEmail implements DatabaseRepository.
func (m *postgresRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, email, username, first_name, last_name, password_hash
		FROM users
		WHERE email = $1
	`

	row := m.db.QueryRowContext(ctx, query, email)

	var user models.User

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserById implements DatabaseRepository.
func (m *postgresRepo) GetUserById(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, email, username, first_name, last_name, password_hash
		FROM users
		WHERE id = $1
	`

	row := m.db.QueryRowContext(ctx, query, id)

	var user models.User

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername implements DatabaseRepository.
func (m *postgresRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, email, username, first_name, last_name, password_hash
		FROM users
		WHERE username = $1
	`

	row := m.db.QueryRowContext(ctx, query, username)

	var user models.User

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateMeeting implements DatabaseRepository.
func (m *postgresRepo) CreateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO meetings (host_id, creator_id, title, description, start_time, end_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var newID int64

	err := m.db.QueryRowContext(
		ctx, query, meeting.HostID, meeting.CreatorID, meeting.Title, meeting.Description, meeting.StartTime, meeting.EndTime,
	).Scan(&newID)
	if err != nil {
		return nil, err
	}

	meeting.ID = newID

	return meeting, nil
}

// DeleteMeeting implements DatabaseRepository.
func (m *postgresRepo) DeleteMeeting(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		DELETE FROM meetings
		WHERE id = $1
	`

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetMeetingById implements DatabaseRepository.
func (m *postgresRepo) GetMeetingById(id int64) (*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, host_id, creator_id, title, descrtiption, start_time, end_time
		FROM meetings
		WHERE id = $1
	`

	row := m.db.QueryRowContext(ctx, query, id)

	var meeting models.Meeting

	if err := row.Scan(
		&meeting.ID, &meeting.HostID, &meeting.CreatorID, &meeting.Title, &meeting.Description, &meeting.StartTime, &meeting.EndTime,
	); err != nil {
		return nil, err
	}

	return &meeting, nil
}

// GetMeetingsForUser implements DatabaseRepository.
func (m *postgresRepo) GetMeetingsForUser(userId int64) ([]*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, host_id, creator_id, title, description, start_time, end_time
		FROM meetings
		WHERE host_id = $1 OR creator_id = $1
	`

	rows, err := m.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []*models.Meeting

	for rows.Next() {
		var meeting models.Meeting
		if err := rows.Scan(
			&meeting.ID, &meeting.HostID, &meeting.CreatorID, &meeting.Title, &meeting.Description, &meeting.StartTime, &meeting.EndTime,
		); err != nil {
			return nil, err
		}
		meetings = append(meetings, &meeting)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return meetings, nil
}

// UpdateMeeting implements DatabaseRepository.
func (m *postgresRepo) UpdateMeeting(meeting *models.Meeting) (*models.Meeting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE meetings
		SET title = $1, description = $2, start_time = $3, end_time = $4
		WHERE id = $5
	`

	_, err := m.db.ExecContext(ctx, query, meeting.Title, meeting.Description, meeting.StartTime, meeting.EndTime, meeting.ID)
	if err != nil {
		return nil, err
	}

	return meeting, nil
}
