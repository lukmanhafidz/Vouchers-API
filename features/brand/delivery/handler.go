package delivery

import (
	"Test/domain"
	"Test/features/common"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type brandHandler struct {
	brandUseCase domain.BrandUseCase
}

func New(bc domain.BrandUseCase) domain.BrandHandler {
	return &brandHandler{
		brandUseCase: bc,
	}
}

// Create implements domain.BrandHandler
func (bh *brandHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newBrand BrandFormat

		token := common.ExtractData(c)
		if token.Role == "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
		}

		bind := c.Bind(&newBrand)
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := bh.brandUseCase.CreateBrand(newBrand.ToModel())

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Register success",
		})
	}
}
