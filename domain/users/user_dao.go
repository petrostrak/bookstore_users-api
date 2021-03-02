package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/logger"
	"bookstore_users-api/utils/errors"
	"fmt"
)

// User data access object (dao) encapsulates the logic
// to persist and retrieve a user object from DB

const (
	errorNoRows                 = "no rows in result set"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, created_at, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE email=? AND password=? AND status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

// Get returns a user with the given ID, or an error
func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying to prepare select query", err)
		return errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	res := stmt.QueryRow(u.ID)
	if getErr := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); getErr != nil {
		logger.Error("error while trying to get user by id", getErr)
		return errors.NewInternalServerError("DB error")
	}

	return nil
}

// Save saves a user to the DB
func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while trying to prepare insert query", err)
		return errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	// alteratively with lower performance
	// users_db.Client.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)
	res, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if saveErr != nil {
		logger.Error("error while trying to save user", saveErr)
		return errors.NewInternalServerError("DB error")
	}

	userID, err := res.LastInsertId()
	if err != nil {
		logger.Error("error while trying to get last inserted id", err)
		return errors.NewInternalServerError("DB error")
	}

	u.ID = userID
	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while trying to prepare update query", err)
		return errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		logger.Error("error while trying to update user", err)
		return errors.NewInternalServerError("DB error")
	}
	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error while trying to prepare delete query", err)
		return errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.ID); err != nil {
		logger.Error("error while trying to delete user", err)
		return errors.NewInternalServerError("DB error")
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error while trying to prepare finduser by status", err)
		return nil, errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error while trying to find user by status", err)
		return nil, errors.NewInternalServerError("DB error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error while trying to scan user row", err)
			return nil, errors.NewInternalServerError("DB error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

// FindByEmailAndPassword returns a user with the given email and password, or an error
func (u *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error while trying to prepare select query", err)
		return errors.NewInternalServerError("DB error")
	}
	defer stmt.Close()

	res := stmt.QueryRow(u.Email, u.Password, StatusActive)
	if getErr := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); getErr != nil {
		logger.Error("error while trying to get user by email and password", getErr)
		return errors.NewInternalServerError("DB error")
	}

	return nil
}
