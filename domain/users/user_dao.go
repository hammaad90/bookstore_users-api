package users

import (
	"fmt"

	"github.com/hammaad90/bookstore_users-api/utils/date_utils"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

// func Get(userId int64) (*User, *errors.RestError) {
// 	return nil, nil
// }

// func Save(user User) *errors.RestError {
// 	return nil
// }

// mock db
var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	currant := userDB[user.Id]
	if currant != nil {
		if currant.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered ", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists ", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	userDB[user.Id] = user
	return nil
}