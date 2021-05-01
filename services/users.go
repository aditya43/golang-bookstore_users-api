package services

import (
	"github.com/aditya43/golang-bookstore_users-api/domain/users"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RESTErr) {
	return &user, nil
}
