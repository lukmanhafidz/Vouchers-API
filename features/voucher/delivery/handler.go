package delivery

import (
	"Test/domain"
	"Test/features/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type voucherHandler struct {
	voucherUseCase domain.VoucherUseCase
}

func New(vc domain.VoucherUseCase) domain.VoucherHandler {
	return &voucherHandler{
		voucherUseCase: vc,
	}
}

// Create implements domain.VoucherHandler
func (vh *voucherHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newVoucher VoucherFormat

		token := common.ExtractData(c)
		if token.Role == "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
		}

		bind := c.Bind(&newVoucher)
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := vh.voucherUseCase.CreateVoucher(newVoucher.ToModel())

		if status == http.StatusBadRequest {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		if status == http.StatusConflict {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"code":    http.StatusConflict,
				"message": "Cant input existing data",
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

// GetAll implements domain.VoucherHandler
func (vh *voucherHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		brandID := c.QueryParam("id")

		id, err := strconv.Atoi(brandID)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		data, status := vh.voucherUseCase.GetAllVoucher(id)

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    http.StatusOK,
			"message": "success update data",
		})
	}
}

// GetByID implements domain.VoucherHandler
func (vh *voucherHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		voucherID := c.QueryParam("id")

		id, err := strconv.Atoi(voucherID)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		data, status := vh.voucherUseCase.GetByIDVoucher(id)

		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    http.StatusOK,
			"message": "success update data",
		})
	}
}
