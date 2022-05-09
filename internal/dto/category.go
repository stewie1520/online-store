package dto

type CreateCategoryDto struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Author   int32  `json:"author"`
	IsActive bool   `json:"isActive"`
}

type FetchManyCategoriesDto struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	Page int32 `json:"page"`
}

func (f *FetchManyCategoriesDto) CalcOffset () int32 {
	if f.Offset != 0 {
		return f.Offset
	}

	if f.Offset == 0 && f.Page == 0 {
		return 0
	}

	if f.Offset == 0 && f.Page != 1 {
		return f.Limit * f.Page - 1
	}

	return 0
}
