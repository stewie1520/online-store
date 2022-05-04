package datastore

import (
	"database/sql"
	"fmt"

	"github.com/stewie1520/online-store/internal/config"
)

func NewPostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.AppConfig.Database.Url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not send ping to database, %w", err)
	}

	return db, nil
}
