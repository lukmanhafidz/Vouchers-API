package usecase

import (
	"Test/config"
	"Test/domain"
	"Test/domain/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Username: "batman", Fullname: "Bruce Wayne", Email: "brucewayne@gmail.com", City: "Jakarta", Password: "polar",
		Phonenumber: "081212121212", Points: 100, Balance: 100000}

	returnData := mockData
	returnData.ID = 1
	returnData.Role = "user"

	invalidData := domain.User{Fullname: "Bruce Wayne", Email: "brucewayne@gmail.com", City: "Jakarta", Password: "polar",
		Phonenumber: "081212121212", Points: 100, Balance: 100000, Role: "user"}

	t.Run("Success register", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(invalidData, config.COST)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST)

		assert.Equal(t, 409, status)
		repo.AssertExpectations(t)
	})

	t.Run("Generate bcrypt error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, 40)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		returnData.Username = ""
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Username: "Jacob", Password: "jacob123"}
	returnData := domain.User{ID: 1, Role: "user", Username: "Jacob", Password: "jacob123"}

	notfound := mockData
	notfound.ID = 0

	t.Run("Succes Login", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$NdU0QGo64Ny6h4hAzyrISu5UPNnE5RCem72ELBUIIKEcClNFv.g4a")
		repo.On("Login", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.Login(mockData)

		assert.Equal(t, 200, status)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything, mock.Anything).Return(domain.User{}).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.Login(mockData)

		assert.Nil(t, res)
		assert.Equal(t, status, 404)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong input", func(t *testing.T) {
		mockData.Password = ""
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8w")
		useCase := New(repo, validator.New())
		res, status := useCase.Login(mockData)

		assert.Nil(t, res)
		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		returnData.ID = 0
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.Login(notfound)

		assert.Equal(t, 500, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
