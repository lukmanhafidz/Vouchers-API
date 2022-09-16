package delivery

import "Test/domain"

type TransactionsFormat struct {
	VoucherID int `json:"voucherID" validate:"required"`
	UserID    int
	Items     int    `json:"items" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

func (f *TransactionsFormat) ToModel() domain.Transaction {
	return domain.Transaction{
		VoucherID: f.VoucherID,
		UserID:    f.UserID,
		Items:     f.Items,
		Status:    f.Status,
	}
}
