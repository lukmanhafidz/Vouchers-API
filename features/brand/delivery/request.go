package delivery

import "Test/domain"

type BrandFormat struct {
	Name string `json:"name" validate:"required"`
}

func (f *BrandFormat) ToModel() domain.Brand {
	return domain.Brand{
		Name: f.Name,
	}
}
