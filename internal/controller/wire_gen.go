// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

import (
	"github.com/stewie1520/internal/infrastructure/datastore"
	"github.com/stewie1520/internal/service"
)

// Injectors from init_category.go:

func InitializeCategoryController() (CategoryController, error) {
	queries, err := datastore.ProvideQueries()
	if err != nil {
		return nil, err
	}
	categoryService := service.NewCategoryService(queries)
	controllerCategoryController := NewCategoryController(categoryService)
	return controllerCategoryController, nil
}