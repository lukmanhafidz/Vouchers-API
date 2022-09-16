package data

import (
	"Test/domain"
)

type Voucher struct {
	Name    string `json:"name" validate:"required"`
	BrandID int    `json:"brandID"`
	Cim     int    `json:"Cim" validate:"required"`
	Cip     int    `json:"Cip" validate:"required"`
}

type Voucher_Brand struct {
	ID        int
	Name      string
	BrandName string
	Cim       int
	Cip       int
	Code      string
}

func (v *Voucher) ToModel() domain.Voucher {
	return domain.Voucher{
		Name:    v.Name,
		BrandID: v.BrandID,
		Cim:     v.Cim,
		Cip:     v.Cip,
	}
}

func (vb *Voucher_Brand) ToModelVoucher_Brand() domain.Voucher_Brand {
	return domain.Voucher_Brand{
		ID:        vb.ID,
		Name:      vb.Name,
		BrandName: vb.BrandName,
		Cim:       vb.Cim,
		Cip:       vb.Cip,
		Code:      vb.Code,
	}
}

func ParseToArr(arr []Voucher_Brand) []domain.Voucher_Brand {
	var res []domain.Voucher_Brand

	for _, val := range arr {
		res = append(res, val.ToModelVoucher_Brand())
	}

	return res
}

func FromModel(data domain.Voucher) Voucher {
	var res Voucher
	res.Name = data.Name
	res.BrandID = data.BrandID
	res.Cim = data.Cim
	res.Cip = data.Cip
	return res
}
