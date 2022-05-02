package service

import (
	"context"

	"github.com/stewie1520/internal/domain"
	"github.com/stewie1520/internal/infrastructure/datastore"
)

func CreateCategory(ctx context.Context, param *domain.CreateCategoryParams) (*domain.Category, error) {
	queries, err := datastore.OpenConnection()
	if err != nil {
		return nil, err
	}

	category, err := queries.CreateCategory(ctx, *param)
	return &category, err
}
