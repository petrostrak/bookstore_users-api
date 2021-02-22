package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"bookstore_users-api/utils/mysql"
)

// User data access object (dao) encapsulates the logic
// to persist and retrieve a user object from DB

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
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
	if getErr := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysql.ParseError(getErr)
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
	res, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if saveErr != nil {
		mysql.ParseError(saveErr)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return mysql.ParseError(err)
	}

	u.ID = userID
	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		return mysql.ParseError(err)
	}
	return nil
}
