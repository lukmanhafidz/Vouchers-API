package data

import (
	"Test/domain"
	"log"

	"gorm.io/gorm"
)

type voucherData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.VoucherData {
	return &voucherData{
		db: db,
	}
}

// CreateVoucherData implements domain.VoucherData
func (vd *voucherData) CreateVoucherData(newVoucher domain.Voucher) domain.Voucher {
	var voucher = FromModel(newVoucher)
	err := vd.db.Create(&voucher).Error
	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.Voucher{}
	}

	return voucher.ToModel()
}

// GetAllVoucherData implements domain.VoucherData
func (vd *voucherData) GetAllVoucherData(brandID int) []domain.Voucher_Brand {
	var voucher []Voucher_Brand

	err := vd.db.Model(&Voucher{}).Select("vouchers.id, vouchers.name, brands.brand_name, vouchers.cim, vouchers.cip").
		Joins("join brands on vouchers.brand_id = brands.id").Scan(&voucher)

	if err.Error != nil {
		log.Println("cant get vouchers data", err.Error.Error())
		return nil
	}

	if err.RowsAffected == 0 {
		log.Println("data not found", err.Error.Error())
		return nil
	}
	log.Println(voucher)
	return ParseToArr(voucher)
}

// GetByIDVoucherData implements domain.VoucherData
func (vd *voucherData) GetByIDVoucherData(id int) domain.Voucher_Brand {
	var voucher Voucher_Brand

	err := vd.db.Model(&Voucher{}).Select("vouchers.id, vouchers.name, brands.brand_name, vouchers.cim, vouchers.cip").
		Joins("join brands on vouchers.brand_id = brands.id").Scan(&voucher)

	if err.Error != nil {
		log.Println("cant get voucher data", err.Error.Error())
		return domain.Voucher_Brand{}
	}

	if err.RowsAffected == 0 {
		log.Println("data not found", err.Error.Error())
		return domain.Voucher_Brand{}
	}

	return voucher.ToModelVoucher_Brand()
}
