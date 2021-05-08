package users

// DTO: Data Transfer Object

import (
	"strings"

	"github.com/aditya43/golang-bookstore_users-api/utils/date_time"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` // To exclude password field when dealing with JSON
}

type Users []User

func (user *User) Validate() *errors.RESTErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.DateCreated = date_time.GetUTCDateTimeAPIFormatDBFormat()
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))

	if user.Email == "" {
		return errors.BadRequestErr("Invalid or empty email address")
	}

	if user.Password == "" {
		return errors.BadRequestErr("Invalid or empty password")
	}

	return nil
}
