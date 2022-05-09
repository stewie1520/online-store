package service

import (
	"context"

	"github.com/stewie1520/online-store/internal/domain"
)

type categoryService struct {
	queries *domain.Queries
}

type CategoryService interface {
	CreateCategory(ctx context.Context, param *domain.CreateCategoryParams) (*domain.Category, error)
	FetchCategories(ctx context.Context, param *domain.FetchCategoriesParams) ([]domain.Category, error)
}

func NewCategoryService(queries *domain.Queries) CategoryService {
	return &categoryService{
		queries: queries,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, param *domain.CreateCategoryParams) (*domain.Category, error) {
	category, err := s.queries.CreateCategory(ctx, *param)
	return &category, err
}

func (s *categoryService) FetchCategories(ctx context.Context, param *domain.FetchCategoriesParams) ([]domain.Category, error) {
	categories, err := s.queries.FetchCategories(ctx, *param)
	return categories, err
}
