package datastore

import (
	"github.com/stewie1520/online-store/internal/domain"
)

func OpenConnection() (*domain.Queries, error) {
	db, err := NewPostgresConnection()
	if err != nil {
		return nil, err
	}

	return domain.New(db), nil
}

func ProvideQueries() (*domain.Queries, error) {
	queries, err := OpenConnection()
	if err != nil {
		return nil, err
	}

	return queries, nil
}
