package users

// DAO: Data Access Object
// This is the only place where we will interact with database

import (
	"fmt"

	"github.com/aditya43/golang-bookstore_users-api/datasources/mysql/users_db"
	"github.com/aditya43/golang-bookstore_users-api/logger"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

const (
	queryInsert       = "INSERT INTO users (`first_name`, `last_name`, `email`, `date_created`, `status`, `password`) VALUES (?, ?, ?, ?, ?, ?)"
	queryGet          = "SELECT `id`, `first_name`, `last_name`, `email`, `date_created`, `status` FROM users WHERE id=?"
	queryUpdate       = "UPDATE users SET `first_name`=?, `last_name`=?, `email`=? WHERE `id`=?"
	queryDelete       = "DELETE FROM users WHERE `id`=?"
	queryFindByStatus = "SELECT `id`, `first_name`, `last_name`, `email`, `date_created`, `status` FROM users WHERE status=?"
)

func (user *User) Get() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryGet)
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}
	defer stmt.Close()

	res := stmt.QueryRow(user.Id)

	if err := res.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}

	return nil
}

func (user *User) Save() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryInsert)
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}
	defer stmt.Close()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}

	userId, err := insertRes.LastInsertId()
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryUpdate)
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
	); err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}

	return nil
}

func (user *User) Delete() *errors.RESTErr {
	stmt, err := users_db.DBClient.Prepare(queryDelete)
	if err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error(err)
		return errors.InternalServerErr("Something went wrong!")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RESTErr) {
	stmt, err := users_db.DBClient.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error(err)
		return nil, errors.InternalServerErr("Something went wrong!")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error(err)
		return nil, errors.InternalServerErr("Something went wrong!")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.DateCreated,
			&user.Status,
		); err != nil {
			logger.Error(err)
			return nil, errors.InternalServerErr("Something went wrong!")
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundErr(fmt.Sprintf("0 results found for status = '%s'", status))
	}

	return results, nil
}
