package usecase

import (
	"Test/domain"
	"Test/features/transaction/data"
	"log"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type transactionUseCase struct {
	transactionData domain.TransactionData
	valid           *validator.Validate
}

func New(td domain.TransactionData, val *validator.Validate) domain.TransactionUseCase {
	return &transactionUseCase{
		transactionData: td,
		valid:           val,
	}
}

// CreateRedeem implements domain.TransactionUseCase
func (tc *transactionUseCase) CreateRedeem(newTrans domain.Transaction, id int) int {
	var transaction = data.FromModel(newTrans)

	validError := tc.valid.Struct(transaction)
	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	if id == 0 {
		log.Println("ID = 0")
		return 500
	}

	price := tc.transactionData.GetVoucherData(transaction.VoucherID)
	transaction.Total = transaction.Items * price
	transaction.Code = uuid.New().String()
	transaction.UserID = id

	create := tc.transactionData.CreateRedeemData(transaction.ToModel())

	if create.VoucherID == 0 {
		log.Println("Empty Data")
		return 500
	}

	return 200
}

// GetRedeem implements domain.TransactionUseCase
func (tc *transactionUseCase) GetRedeem(id int) (map[string]interface{}, int) {
	var res = map[string]interface{}{}

	data := tc.transactionData.GetRedeemData(id)
	if data.VoucherID == 0 {
		return nil, 404
	}

	res["voucherID"] = data.VoucherID
	res["voucherName"] = data.VoucherName
	res["brandName"] = data.BrandName
	res["cim"] = data.Cim
	res["cip"] = data.Cip
	res["code"] = data.Code
	res["userID"] = data.UserID
	res["username"] = data.Username
	res["email"] = data.Email
	res["items"] = data.Items
	res["status"] = data.Status
	res["total"] = data.Total

	return res, 200
}
