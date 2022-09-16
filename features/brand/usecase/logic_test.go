package usecase

import (
	"Test/domain"
	"Test/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBrand(t *testing.T) {
	repo := new(mocks.BrandData)

	mockData := domain.Brand{BrandName: "indomaret"}
	returnData := domain.Brand{ID: 1, BrandName: "indomaret"}

	t.Run("Success Create", func(t *testing.T) {
		repo.On("CreateBrandData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateBrand(mockData)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.CreateBrand(domain.Brand{})

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		repo.On("CreateBrandData", mock.Anything).Return(domain.Brand{}).Once()
		useCase := New(repo, validator.New())
		status := useCase.CreateBrand(mockData)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})
}
