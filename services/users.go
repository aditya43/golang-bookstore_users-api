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

func UpdateUser(isPatchMethod bool, user users.User) (*users.User, *errors.RESTErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if !isPatchMethod {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	} else {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}

		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}

		if user.Email != "" {
			currentUser.Email = user.Email
		}
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userId int64) *errors.RESTErr {
	user, err := GetUser(userId)
	if err != nil {
		return err
	}

	user.Id = userId
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
