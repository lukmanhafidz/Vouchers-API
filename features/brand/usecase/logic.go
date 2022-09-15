package usecase

import (
	"Test/domain"
	"Test/features/brand/data"
	"log"

	"github.com/go-playground/validator"
)

type brandUseCase struct {
	brandData domain.BrandData
	valid     *validator.Validate
}

func New(bd domain.BrandData, val *validator.Validate) domain.BrandUseCase {
	return &brandUseCase{
		brandData: bd,
		valid:     val,
	}
}

// CreateBrand implements domain.BrandUseCase
func (bc *brandUseCase) CreateBrand(newBrand domain.Brand) int {
	var brand = data.FromModel(newBrand)

	validError := bc.valid.Struct(brand)
	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	create := bc.brandData.CreateBrandData(brand.ToModel())
	if create.Name == "" {
		log.Println("Empty Data")
		return 500
	}

	return 200
}
