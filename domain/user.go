package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID           int
	Username     string
	Email        string
	City         string
	PhoneNumber  string
	Password     string
	Points       int
	Balance      int
	Role         string
	Transactions []Transactions
}

type UserHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}

type UserUseCase interface {
	Login(userdata User) (User, error)
	RegisterUser(newuser User, cost int) int
}

type UserData interface {
	Login(userdata User) User
	RegisterData(newuser User) User
	CheckDuplicate(newuser User) bool
	GetPasswordData(name string) string
}
