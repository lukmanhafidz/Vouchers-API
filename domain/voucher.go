package domain

import "github.com/labstack/echo/v4"

type Voucher struct {
	ID      int
	Name    string
	BrandID int
	Cim     int
	Cip     int
	Code    string
}

type Voucher_Brand struct {
	ID        int
	Name      string
	BrandName string
	Cim       int
	Cip       int
	Code      string
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
