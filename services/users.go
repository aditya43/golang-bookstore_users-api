package services

import (
	"github.com/aditya43/golang-bookstore_users-api/domain/users"
	"github.com/aditya43/golang-bookstore_users-api/utils/crypto_utils"
	"github.com/aditya43/golang-bookstore_users-api/utils/date_time"
	"github.com/aditya43/golang-bookstore_users-api/utils/errors"
)

var UserService userServiceInterface = &userService{}

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RESTErr)
	Authenticate(*users.LoginRequest) (*users.User, *errors.RESTErr)
	CreateUser(users.User) (*users.User, *errors.RESTErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RESTErr)
	DeleteUser(int64) *errors.RESTErr
	Search(string) (users.Users, *errors.RESTErr)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RESTErr) {
	user := &users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Authenticate(loginRequest *users.LoginRequest) (*users.User, *errors.RESTErr) {
	user := &users.User{
		Email:    loginRequest.Email,
		Password: crypto_utils.GetMD5(loginRequest.Password),
	}

	if err := user.FindByEmailPassword(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RESTErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_time.GetUTCDateTimeAPIFormatDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(isPatchMethod bool, user users.User) (*users.User, *errors.RESTErr) {
	currentUser, err := UserService.GetUser(user.Id)
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

func (s *userService) DeleteUser(userId int64) *errors.RESTErr {
	user, err := UserService.GetUser(userId)
	if err != nil {
		return err
	}

	user.Id = userId
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (s *userService) Search(status string) (users.Users, *errors.RESTErr) {
	return (&users.User{}).FindByStatus(status)
}
