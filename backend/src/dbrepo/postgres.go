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
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var newID int64

	if err := m.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password).Scan(&newID); err != nil {
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
		SELECT id, email, username, password_hash
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
		SELECT id, email, username, password_hash
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
