package data

import (
	"Test/domain"
	"log"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

// CheckDuplicate implements domain.UserData
func (ud *userData) CheckDuplicate(newUser domain.User) bool {
	var user = FromModel(newUser)

	err := ud.db.Find(&user, "username = ? OR email = ?", user.Username, user.Email)
	if err.RowsAffected == 1 {
		log.Println("Duplicated data", err.Error)
		return true
	}

	return false
}

// GetPasswordData implements domain.UserData
func (ud *userData) GetPasswordData(name string) string {
	var user User
	err := ud.db.Find(&user, "username = ?", name).Error

	if err != nil {
		log.Println("Cant retrieve user data", err.Error())
		return ""
	}

	return user.Password
}

// Login implements domain.UserData
func (ud *userData) Login(userData domain.User) domain.User {
	var user = FromModel(userData)
	var err error

	err = ud.db.First(&user, "username = ?", userData.Username).Error

	if err != nil {
		log.Println("Cant login data", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

// RegisterData implements domain.UserData
func (ud *userData) RegisterData(newUser domain.User) domain.User {
	var user = FromModel(newUser)
	err := ud.db.Create(&user).Error

	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}
