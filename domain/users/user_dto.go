package users

import "github.com/hammaad90/bookstore_users-api/utils/errors"

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// a function
// func Validate(user *User) *errors.RestError {
// 	if  user.Email == ""{
// 		return nil, errors.NewBadRequestError(message: "invalid email address")
// 	}

// }

// method
func (user *User) Validate() *errors.RestError {
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
