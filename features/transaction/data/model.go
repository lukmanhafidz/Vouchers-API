package data

import "Test/domain"

type Transaction struct {
	VoucherID int `json:"voucherID" validate:"required"`
	UserID    int
	Items     int    `json:"items" validate:"required"`
	Status    string `json:"status" validate:"required"`
	Code      string
	Total     int
}

type Transaction_Junction struct {
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

func (t *Transaction) ToModel() domain.Transaction {
	return domain.Transaction{
		VoucherID: t.VoucherID,
		UserID:    t.UserID,
		Items:     t.Items,
		Status:    t.Status,
		Code:      t.Code,
		Total:     t.Total,
	}
}

func (tj *Transaction_Junction) ToModelTransactions_Junction() domain.Transaction_Junction {
	return domain.Transaction_Junction{
		VoucherID:   tj.VoucherID,
		VoucherName: tj.VoucherName,
		BrandName:   tj.BrandName,
		Cim:         tj.Cim,
		Cip:         tj.Cip,
		Code:        tj.Code,
		UserID:      tj.UserID,
		Username:    tj.Username,
		Email:       tj.Email,
		Items:       tj.Items,
		Status:      tj.Status,
		Total:       tj.Total,
	}
}

func FromModel(data domain.Transaction) Transaction {
	var res Transaction
	res.VoucherID = data.VoucherID
	res.UserID = data.UserID
	res.Items = data.Items
	res.Status = data.Status
	res.Code = data.Code
	res.Total = data.Total
	return res
}
