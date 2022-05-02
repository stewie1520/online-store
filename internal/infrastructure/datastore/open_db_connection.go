package datastore

import (
	"github.com/stewie1520/internal/domain"
)

func OpenConnection() (*domain.Queries, error) {
	db, err := OpenPostgresConnection()
	if err != nil {
		return nil, err
	}

	return domain.New(db), nil
}