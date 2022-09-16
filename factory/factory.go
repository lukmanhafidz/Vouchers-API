package factory

import (
	ud "Test/features/user/data"
	udeli "Test/features/user/delivery"
	uc "Test/features/user/usecase"

	bd "Test/features/brand/data"
	bdeli "Test/features/brand/delivery"
	bc "Test/features/brand/usecase"

	vd "Test/features/voucher/data"
	vdeli "Test/features/voucher/delivery"
	vc "Test/features/voucher/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	valid := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, valid)
	userHandler := udeli.New(userCase)
	udeli.RouteUser(e, userHandler)

	brandData := bd.New(db)
	brandCase := bc.New(brandData, valid)
	brandHandler := bdeli.New(brandCase)
	bdeli.RouteBrand(e, brandHandler)

	voucherData := vd.New(db)
	voucherCase := vc.New(voucherData, valid)
	voucherHandler := vdeli.New(voucherCase)
	vdeli.RouteVoucher(e, voucherHandler)
}
