package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

func (db *DB) Close() error {
	return db.SQL.Close()
}

const (
	maxOpenDBConns    = 1
	maxIdleDBConns    = 1
	maxDBConnLifetime = 1 * time.Minute
)

func ConnectSQL(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConns)
	db.SetMaxIdleConns(maxIdleDBConns)
	db.SetConnMaxLifetime(maxDBConnLifetime)

	d := &DB{
		SQL: db,
	}

	return d, nil
}
