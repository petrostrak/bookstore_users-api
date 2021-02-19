package users

import (
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"fmt"
)

// User data access object (dao) encapsulates the logic
// to persist and retrieve a user object from DB

var (
	usersDB = make(map[int64]*User)
)

// Get returns a user with the given ID, or an error
func (u *User) Get() *errors.RestErr {
	result := usersDB[u.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", result.ID))
	}
	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated

	return nil
}

// Save saves a user to the DB
func (u *User) Save() *errors.RestErr {
	current := usersDB[u.ID]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.ID))
	}

	u.DateCreated = date.GetNowString()

	usersDB[u.ID] = u
	return nil
}
