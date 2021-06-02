package services

import (
	"github.com/hammaad90/bookstore_users-api/domain/users"
	"github.com/hammaad90/bookstore_users-api/utils/crypto_utils"
	"github.com/hammaad90/bookstore_users-api/utils/date_utils"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

var UserService usersServiceInterface = &userService{}

type userService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	Search(string) (users.Users, *errors.RestError)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	// we always return either result or error nor both we never return nil, nil
	// using pointer coz other wise we have to allocate it to variable and that is a waste of memory, thats why we are using pointer
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDbFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.FirstName != "" {
			current.LastName = user.LastName
		}

		if user.FirstName != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

func (s *userService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}
