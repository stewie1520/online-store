//go:build wireinject
// +build wireinject

package controller

import (
	"github.com/google/wire"
	"github.com/stewie1520/online-store/internal/infrastructure/datastore"
	"github.com/stewie1520/online-store/internal/service"
)

func NewCategoryController() (CategoryController, error) {
	wire.Build(
		newCategoryController,
		service.NewCategoryService,
		datastore.ProvideQueries,
	)

	return nil, nil
}
