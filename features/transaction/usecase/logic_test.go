package usecase

import (
	"Test/domain"
	"Test/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRedeem(t *testing.T) {
	repo := new(mocks.TransactionData)

	mockData := domain.Transaction{VoucherID: 1, UserID: 1, Items: 2, Status: "ok"}

	returnData := mockData
	returnData.ID = 1
	returnData.Total = 20
	returnData.Code = "12jnp-sdni1-00sa2"

	invalidData := domain.Transaction{ID: 0, UserID: 1, Items: 2, Status: "ok"}

	t.Run("Success Create", func(t *testing.T) {
		repo.On("GetVoucherData", mock.Anything).Return(1).Once()
		repo.On("CreateRedeemData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateRedeem(mockData, 1)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.CreateRedeem(invalidData, 1)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("User ID Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.CreateRedeem(mockData, 0)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		repo.On("GetVoucherData", mock.Anything).Return(1).Once()
		repo.On("CreateRedeemData", mock.Anything).Return(invalidData).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateRedeem(mockData, 1)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})
}

func TestGetRedeem(t *testing.T) {
	repo := new(mocks.TransactionData)

	returnData := domain.Transaction_Junction{ID: 1, VoucherID: 1, UserID: 1, Items: 2, Status: "ok", VoucherName: "50 discount",
		BrandName: "indomaret", Cim: 1000, Cip: 10, Code: "12jnp-sdni1-00sa2", Username: "Jacob", Email: "jacob@gmail.com", Total: 20}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("GetRedeemData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetRedeem(1)

		assert.Equal(t, 200, status)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetRedeemData", mock.Anything).Return(domain.Transaction_Junction{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetRedeem(1)

		assert.Equal(t, 404, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
