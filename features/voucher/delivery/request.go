package delivery

import "Test/domain"

type VoucherFormat struct {
	Name    string `json:"name" validate:"required"`
	BrandID int    `json:"brandID"`
	Cim     int    `json:"cim" validate:"required"`
	Cip     int    `json:"cip" validate:"required"`
}

func (f *VoucherFormat) ToModel() domain.Voucher {
	return domain.Voucher{
		Name:    f.Name,
		BrandID: f.BrandID,
		Cim:     f.Cim,
		Cip:     f.Cip,
	}
}
