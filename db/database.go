package db

import (
	"context"
	"database/sql"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(databasePath string) (*Database, error) {
	db, err := sql.Open("sqlite", databasePath)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}


