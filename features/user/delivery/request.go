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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: i.Username,
		Password: i.Password,
	}
}

func (i *UserFormat) ToModel() domain.User {
	return domain.User{
		Username:    i.Username,
		Email:       i.Email,
		City:        i.City,
		Phonenumber: i.Phonenumber,
		Password:    i.Password,
		Points:      i.Points,
		Balance:     i.Balance,
	}
}
