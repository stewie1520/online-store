//go:build wireinject
// +build wireinject

package controller

import (
	"github.com/google/wire"
	"github.com/stewie1520/internal/infrastructure/datastore"
	"github.com/stewie1520/internal/service"
)

func InitializeCategoryController() (CategoryController, error) {
	wire.Build(
		NewCategoryController,
		service.NewCategoryService,
		datastore.ProvideQueries,
	)

	return nil, nil
}
