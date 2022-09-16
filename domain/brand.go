package domain

import "github.com/labstack/echo/v4"

type Brand struct {
	ID        int
	BrandName string
}

type BrandHandler interface {
	Create() echo.HandlerFunc
}

type BrandUseCase interface {
	CreateBrand(newBrand Brand) int
}

type BrandData interface {
	CreateBrandData(newBrand Brand) Brand
}
