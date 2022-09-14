package domain

import "github.com/labstack/echo/v4"

type Voucher struct {
	ID           int
	Name         string
	BrandID      int
	CiM          int
	CiP          int
	Code         string
	Transactions []Transactions
}

type VoucherHandler interface {
	Create() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type VoucherUseCase interface {
	CreateVoucher(newVoucher Voucher) int
	GetByIDVoucher(id int) int
	GetAllVoucher(brandID int) int
}

type VoucherData interface {
	CreateVoucherData(newVoucher Voucher) Voucher
	GetByIDVoucherData(id int) Voucher
	GetAllVoucherData(brandID int) []Voucher
}
