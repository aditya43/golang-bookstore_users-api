package users

// DAO: Data Access Object
// This is the only place where we will interact with database

import (
	"fmt"

	"github.com/aditya43/golang-bookstore_users-api/datasources/mysql/users_db"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users (`first_name`, `last_name`, `email`, `date_created`) VALUES (?, ?, ?, ?)"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RESTErr {
	if err := users_db.DBClient.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := users_db.DBClient.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}

	userId, err := insertRes.LastInsertId()
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}

	user.Id = userId
	return nil
}
