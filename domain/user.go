package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID          int
	Username    string
	Fullname    string
	Email       string
	City        string
	Phonenumber string
	Password    string
	Points      int
	Balance     int
	Role        string
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}

type UserUseCase interface {
	Login(userData User) (map[string]interface{}, int)
	RegisterUser(newUser User, cost int) int
}

type UserData interface {
	Login(userLogin User) User
	RegisterData(newUser User) User
	CheckDuplicate(newUser User) bool
	GetPasswordData(name string) string
}
