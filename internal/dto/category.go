package dto

type CreateCategoryDto struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Author   int32  `json:"author"`
	IsActive bool   `json:"isActive"`
}
