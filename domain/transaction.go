package domain

import "github.com/labstack/echo/v4"

type Transactions struct {
	ID        int
	VoucherID int
	UserID    int
	Status    string
}

type TransactionsHandler interface {
	Create() echo.HandlerFunc
	Get() echo.HandlerFunc
}

type TransactionsUseCase interface {
	CreateRedeem(newTrans Transactions) int
	GetRedeem(id int) int
}

type TransactionsData interface {
	CreateRedeemData(newTrans Transactions) Transactions
	GetRedeemData(id int) Transactions
}
