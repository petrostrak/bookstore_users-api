package users_db

import (
	"bookstore_users-api/env"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = os.Getenv(env.My_sql_users_username)
	password = os.Getenv(env.My_sql_users_password)
	host     = os.Getenv(env.My_sql_users_host)
	schema   = os.Getenv(env.My_sql_users_schema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema)
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
