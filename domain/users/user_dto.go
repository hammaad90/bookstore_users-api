package users

import (
	"strings"

	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

// a function
// func Validate(user *User) *errors.RestError {
// 	if  user.Email == ""{
// 		return nil, errors.NewBadRequestError(message: "invalid email address")
// 	}

// }

// method
//we are assigning a validate method to user struct, so in this way user know how to validate itself and return whatever error we got
// why are we taking a pointer to user and why are we returning a pointer to error?
func (user *User) Validate() *errors.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password !!! ")
	}
	return nil
}
