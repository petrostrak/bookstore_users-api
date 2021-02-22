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
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?;"
)

var (
	usersDB = make(map[int64]*User)
)

// Get returns a user with the given ID, or an error
func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(u.ID)
	if err := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to get user %d: %s", u.ID, err.Error()))
	}

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
