package app

import "bookstore_users-api/controllers/users"

func mapURLs() {
	router.GET("/hello", users.Hello)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.POST("/users", users.CreateUser)
}
