package factory

import (
	ud "Test/features/user/data"
	udeli "Test/features/user/delivery"
	uc "Test/features/user/usecase"

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
}
