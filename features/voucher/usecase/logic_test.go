package usecase

import (
	"Test/domain"
	"Test/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateVoucher(t *testing.T) {
	repo := new(mocks.VoucherData)

	mockData := domain.Voucher{VoucherName: "50 discount", BrandID: 1, Cim: 1000, Cip: 10}

	returnData := mockData
	returnData.ID = 1

	invalidData := mockData
	invalidData.VoucherName = ""

	t.Run("Success Create", func(t *testing.T) {
		repo.On("CreateVoucherData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateVoucher(mockData)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.CreateVoucher(invalidData)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		repo.On("CreateVoucherData", mock.Anything).Return(invalidData).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateVoucher(mockData)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})
}

func TestGetAllVoucher(t *testing.T) {
	repo := new(mocks.VoucherData)

	mockData := []domain.Voucher_Brand{{VoucherName: "50 discount", BrandName: "indomaret", Cim: 1000, Cip: 10}}

	returnData := mockData
	returnData[0].ID = 1

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("GetAllVoucherData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetAllVoucher(1)

		assert.Equal(t, 200, status)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAllVoucherData", mock.Anything).Return([]domain.Voucher_Brand{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetAllVoucher(1)

		assert.Equal(t, 404, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

}

func TestGetVoucherByID(t *testing.T) {
	repo := new(mocks.VoucherData)

	mockData := domain.Voucher_Brand{VoucherName: "50 discount", BrandName: "indomaret", Cim: 1000, Cip: 10}

	returnData := mockData
	returnData.ID = 1

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("GetByIDVoucherData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetByIDVoucher(1)

		assert.Equal(t, 200, status)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetByIDVoucherData", mock.Anything).Return(domain.Voucher_Brand{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetByIDVoucher(1)

		assert.Equal(t, 404, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

}
