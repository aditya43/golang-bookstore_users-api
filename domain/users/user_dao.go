package users

// DAO: Data Access Object
// This is the only place where we will interact with database

import (
	"github.com/aditya43/golang-bookstore_users-api/datasources/mysql/users_db"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users (`first_name`, `last_name`, `email`, `date_created`) VALUES (?, ?, ?, ?)"
	queryGetUser    = "SELECT `id`, `first_name`, `last_name`, `email`, `date_created` FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET `first_name`=?, `last_name`=?, `email`=? WHERE `id`=?"
	queryDeleteUser = "DELETE FROM users WHERE `id`=?"
)

func (user *User) Get() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryGetUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(user.Id)

	if err := res.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
	); err != nil {
		return errors.InternalServerErr(err.Error())
	}

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

func (user *User) Update() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
	); err != nil {
		return errors.InternalServerErr(err.Error())
	}

	return nil
}

func (user *User) Delete() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return errors.InternalServerErr(err.Error())
	}

	return nil
}
