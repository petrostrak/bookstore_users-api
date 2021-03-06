package users

import (
	"bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

// User data transfer object (dto) is the object that we are transfering
// from the persistance layer, to the application and backwards
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"created_at"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
