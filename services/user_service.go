package services

import (
	"github.com/hammaad90/bookstore_users-api/domain/users"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	// we always return either result or error nor both we never return nil, nil
	// using pointer coz other wise we have to allocate it to variable and that is a waste of memory, thats why we are using pointer
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
