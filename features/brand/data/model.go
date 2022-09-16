package data

import "Test/domain"

type Brand struct {
	BrandName string `json:"name" validate:"required"`
}

func (b *Brand) ToModel() domain.Brand {
	return domain.Brand{
		BrandName: b.BrandName,
	}
}

func FromModel(data domain.Brand) Brand {
	var res Brand
	res.BrandName = data.BrandName
	return res
}
