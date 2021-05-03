package users

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"-"` // To exclude password field when dealing with JSON
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"` // To exclude password field when dealing with JSON
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	_ = json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))

	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}

	return result
}
