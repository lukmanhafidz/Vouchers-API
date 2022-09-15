package delivery

import "Test/domain"

type UserFormat struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email"`
	City        string `json:"city"`
	Phonenumber string `json:"phoneNumber"`
	Password    string `json:"password" validate:"required"`
	Points      int    `json:"points"`
	Balance     int    `json:"balance"`
}
type LoginFormat struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (f *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: f.Username,
		Password: f.Password,
	}
}

func (f *UserFormat) ToModel() domain.User {
	return domain.User{
		Username:    f.Username,
		Email:       f.Email,
		City:        f.City,
		Phonenumber: f.Phonenumber,
		Password:    f.Password,
		Points:      f.Points,
		Balance:     f.Balance,
	}
}
