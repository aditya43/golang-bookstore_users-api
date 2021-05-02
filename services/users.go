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

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RESTErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}
