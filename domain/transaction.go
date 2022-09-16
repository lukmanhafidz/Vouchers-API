package domain

import "github.com/labstack/echo/v4"

type Transaction struct {
	ID        int
	VoucherID int
	UserID    int
	Items     int
	Status    string
	Code      string
	Total     int
}

type Transaction_Junction struct {
	ID          int
	VoucherID   int
	VoucherName string
	BrandName   string
	Cim         int
	Cip         int
	Code        string
	UserID      int
	Username    string
	Email       string
	Items       int
	Status      string
	Total       int
}

type TransactionHandler interface {
	Create() echo.HandlerFunc
	Get() echo.HandlerFunc
}

type TransactionUseCase interface {
	CreateRedeem(newTrans Transaction, id int) int
	GetRedeem(id int) (map[string]interface{}, int)
}

type TransactionData interface {
	CreateRedeemData(newTrans Transaction) Transaction
	GetRedeemData(id int) Transaction_Junction
	GetVoucherData(id int) int
}
