package usecase

import (
	"Test/domain"
	"Test/features/voucher/data"
	"log"

	"github.com/go-playground/validator"
)

type voucherUseCase struct {
	voucherData domain.VoucherData
	valid       *validator.Validate
}

func New(vd domain.VoucherData, val *validator.Validate) domain.VoucherUseCase {
	return &voucherUseCase{
		voucherData: vd,
		valid:       val,
	}
}

// CreateVoucher implements domain.VoucherUseCase
func (vc *voucherUseCase) CreateVoucher(newVoucher domain.Voucher) int {
	var voucher = data.FromModel(newVoucher)

	validError := vc.valid.Struct(voucher)
	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	create := vc.voucherData.CreateVoucherData(voucher.ToModel())

	if create.VoucherName == "" {
		log.Println("Empty Data")
		return 500
	}

	return 200
}

// GetAllVoucher implements domain.VoucherUseCase
func (vc *voucherUseCase) GetAllVoucher(brandID int) ([]map[string]interface{}, int) {
	var arrmap = []map[string]interface{}{}

	data := vc.voucherData.GetAllVoucherData(brandID)
	if len(data) == 0 {
		return nil, 404
	}

	for i := 0; i < len(data); i++ {
		var res = map[string]interface{}{}
		res["voucherID"] = data[i].ID
		res["name"] = data[i].VoucherName
		res["brandName"] = data[i].BrandName
		res["cim"] = data[i].Cim
		res["cip"] = data[i].Cip

		arrmap = append(arrmap, res)
	}

	return arrmap, 200
}

// GetByIDVoucher implements domain.VoucherUseCase
func (vc *voucherUseCase) GetByIDVoucher(id int) (map[string]interface{}, int) {
	var res = map[string]interface{}{}

	data := vc.voucherData.GetByIDVoucherData(id)
	if data.ID == 0 {
		return nil, 404
	}

	res["voucherID"] = data.ID
	res["name"] = data.VoucherName
	res["brandName"] = data.BrandName
	res["cim"] = data.Cim
	res["cip"] = data.Cip

	return res, 200
}
