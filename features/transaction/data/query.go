package data

import (
	"Test/domain"
	"log"

	"gorm.io/gorm"
)

type transactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.TransactionData {
	return &transactionData{
		db: db,
	}
}

// GetVoucherData implements domain.TransactionData
func (td *transactionData) GetVoucherData(id int) int {
	var price int
	err := td.db.Model(&Transaction{}).Select("vouchers.cip").Joins("join vouchers on transactions.voucher_id = vouchers.id").
		Where("transactions.voucher_id = ?", id).Scan(&price)

	if err.Error != nil {
		log.Println("cant get voucher data", err.Error.Error())
		return 0
	}

	if err.RowsAffected == 0 {
		log.Println("data not found", err.Error.Error())
		return 0
	}

	return price
}

// CreateRedeemData implements domain.TransactionsData
func (td *transactionData) CreateRedeemData(newTrans domain.Transaction) domain.Transaction {
	var transaction = FromModel(newTrans)
	var data domain.Transaction

	err := td.db.Create(&transaction).Error
	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.Transaction{}
	}

	td.db.Last(&transaction, "transactions.user_id = ?", transaction.UserID).Scan(&data)

	td.db.Exec("update users join transactions on transactions.user_id = users.id set users.points = users.points - transactions.total where transactions.id = ?", data.ID)

	return transaction.ToModel()
}

// GetRedeemData implements domain.TransactionsData
func (td *transactionData) GetRedeemData(id int) domain.Transaction_Junction {
	var transaction Transaction_Junction

	err := td.db.Model(&Transaction{}).Select("transactions.voucher_id ,vouchers.voucher_name ,brands.brand_name ,vouchers.cim ,vouchers.cip ,transactions.code ,transactions.user_id ,users.username ,users.email ,transactions.items ,transactions.status, transactions.total").
		Joins("join vouchers on transactions.voucher_id = vouchers.id").Joins("join brands on vouchers.brand_id = brands.id").
		Joins("join users on transactions.user_id = users.id").Where("transactions.id = ?", id).Scan(&transaction)

	if err.Error != nil {
		log.Println("cant get transaction data", err.Error.Error())
		return domain.Transaction_Junction{}
	}

	if err.RowsAffected == 0 {
		log.Println("data not found", err.Error.Error())
		return domain.Transaction_Junction{}
	}

	return transaction.ToModelTransactions_Junction()
}
