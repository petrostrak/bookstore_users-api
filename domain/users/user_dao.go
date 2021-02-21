package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

// User data access object (dao) encapsulates the logic
// to persist and retrieve a user object from DB

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

// Get returns a user with the given ID, or an error
func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date.GetNowString()

	// alteratively with lower performance
	// users_db.Client.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)
	res, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintln("error trying to save user", err.Error()))
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error while trying to save user")
	}

	u.ID = userID
	return nil
}
