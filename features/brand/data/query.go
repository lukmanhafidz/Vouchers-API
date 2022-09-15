package data

import (
	"Test/domain"
	"log"

	"gorm.io/gorm"
)

type brandData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.BrandData {
	return &brandData{
		db: db,
	}
}

// CreateBrandData implements domain.BrandData
func (bd *brandData) CreateBrandData(newBrand domain.Brand) domain.Brand {
	var brand = FromModel(newBrand)
	err := bd.db.Create(&brand).Error

	if err != nil {
		log.Println("Cant create brand object", err.Error())
		return domain.Brand{}
	}

	return brand.ToModel()
}
