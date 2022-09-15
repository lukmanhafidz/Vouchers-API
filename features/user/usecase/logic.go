package usecase

import (
	"Test/domain"
	"Test/features/common"
	"Test/features/user/data"
	"log"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
	valid    *validator.Validate
}

func New(ud domain.UserData, val *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		valid:    val,
	}
}

// Login implements domain.UserUseCase
func (uc *userUseCase) Login(userLogin domain.User) (map[string]interface{}, int) {
	var resMap = map[string]interface{}{}

	hashPw := uc.userData.GetPasswordData(userLogin.Username)
	err := bcrypt.CompareHashAndPassword([]byte(hashPw), []byte(userLogin.Password))
	if err != nil {
		log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
		return nil, 400
	}

	login := uc.userData.Login(userLogin)
	if login.Password == "" {
		log.Println("Data login not found")
		return nil, 404
	}

	tokenjwt := common.GenerateToken(login)

	resMap["token"] = tokenjwt
	resMap["username"] = login.Username
	resMap["role"] = login.Role

	return resMap, 200
}

// RegisterUser implements domain.UserUseCase
func (uc *userUseCase) RegisterUser(newUser domain.User, cost int) int {
	var user = data.FromModel(newUser)

	validError := uc.valid.Struct(user)
	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	duplicate := uc.userData.CheckDuplicate(user.ToModel())
	if duplicate {
		log.Println("Duplicate Data")
		return 409
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return 500
	}

	user.Password = string(hashed)
	user.Role = "user"
	user.Balance = 0
	user.Points = 0
	insert := uc.userData.RegisterData(user.ToModel())

	if insert.Username == "" {
		log.Println("Empty Data")
		return 500
	}

	return 200
}
