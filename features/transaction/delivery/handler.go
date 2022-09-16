package delivery

import (
	"Test/domain"
	"Test/features/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionUseCase domain.TransactionUseCase
}

func New(tc domain.TransactionUseCase) domain.TransactionHandler {
	return &transactionHandler{
		transactionUseCase: tc,
	}
}

// Create implements domain.TransactionHandler
func (th *transactionHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newTrans TransactionsFormat
		token := common.ExtractData(c)

		bind := c.Bind(&newTrans)
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := th.transactionUseCase.CreateRedeem(newTrans.ToModel(), token.ID)

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

// Get implements domain.TransactionHandler
func (th *transactionHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID := c.QueryParam("transactionId")

		id, err := strconv.Atoi(transactionID)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		data, status := th.transactionUseCase.GetRedeem(id)

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
