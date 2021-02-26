package users_db

import (
	"bookstore_users-api/env"
	"database/sql"
	"fmt"
	"log"

	// github.com/go-sql-driver/mysql no need to be initialized
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Client is responsible for the connection to DB
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		env.MySQLUsersUsername,
		env.MySQLUsersPassword,
		env.MySQLUsersHost,
		env.MySQLUsersSchema)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB successfully connected")
}
