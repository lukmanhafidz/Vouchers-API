package delivery

import "Test/domain"

type BrandFormat struct {
	BrandName string `json:"name" validate:"required"`
}

func (f *BrandFormat) ToModel() domain.Brand {
	return domain.Brand{
		BrandName: f.BrandName,
	}
}
