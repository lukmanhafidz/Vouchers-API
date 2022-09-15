package delivery

import (
	"Test/config"
	"Test/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUseCase domain.UserUseCase
}

func New(us domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		userUseCase: us,
	}
}

// Login implements domain.UserHandler
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat

		errLog := c.Bind(&userLogin)
		if errLog != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "There is an error in internal server",
			})
		}

		data, status := uh.userUseCase.Login(userLogin.ToModelLogin())
		if status == 400 {
			log.Println("Login failed")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong username or password",
			})
		}

		if status == 404 {
			log.Println("Login failed")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Wrong username or password",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    http.StatusOK,
			"message": "Login success",
		})
	}
}

// Register implements domain.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser UserFormat

		bind := c.Bind(&newUser)
		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Wrong input",
			})
		}

		status := uh.userUseCase.RegisterUser(newUser.ToModel(), config.COST)

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
