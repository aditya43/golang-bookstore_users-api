package services

import (
	"github.com/aditya43/golang-bookstore_users-api/domain/users"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RESTErr) {
	user := &users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RESTErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}
