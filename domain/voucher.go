package domain

import "github.com/labstack/echo/v4"

type Voucher struct {
	ID          int
	VoucherName string
	BrandID     int
	Cim         int
	Cip         int
}

type Voucher_Brand struct {
	ID          int
	VoucherName string
	BrandName   string
	Cim         int
	Cip         int
}

type VoucherHandler interface {
	Create() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type VoucherUseCase interface {
	CreateVoucher(newVoucher Voucher) int
	GetByIDVoucher(id int) (map[string]interface{}, int)
	GetAllVoucher(brandID int) ([]map[string]interface{}, int)
}

type VoucherData interface {
	CreateVoucherData(newVoucher Voucher) Voucher
	GetByIDVoucherData(id int) Voucher_Brand
	GetAllVoucherData(brandID int) []Voucher_Brand
}
