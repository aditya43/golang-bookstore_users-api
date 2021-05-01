package users

// DAO: Data Access Object
// This is the only place where we will interact with database

import (
	"fmt"

	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RESTErr {
	res := usersDB[user.Id]

	if res == nil {
		return errors.NotFoundErr(fmt.Sprintf("User Id %d not found", user.Id))
	}

	user.Id = res.Id
	user.Email = res.Email
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.DateCreated = res.DateCreated

	return nil
}

func (user *User) Save() *errors.RESTErr {
	inDBUser := usersDB[user.Id]

	if inDBUser != nil {
		if inDBUser.Email == user.Email {
			return errors.BadRequestErr(fmt.Sprintf("User Email '%s' already exists", user.Email))
		}

		return errors.BadRequestErr(fmt.Sprintf("User with id %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
