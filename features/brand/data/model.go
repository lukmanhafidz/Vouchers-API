package data

import "Test/domain"

type Brand struct {
	Name string `json:"name" validate:"required"`
}

func (b *Brand) ToModel() domain.Brand {
	return domain.Brand{
		Name: b.Name,
	}
}

func FromModel(data domain.Brand) Brand {
	var res Brand
	res.Name = data.Name
	return res
}
