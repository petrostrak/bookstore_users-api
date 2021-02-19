package app

import "bookstore_users-api/controllers/users"

func mapURLs() {
	// curl -X GET localhost:8000/users/123 -v
	router.GET("/hello", users.Hello)

	// curl -X GET localhost:8000/users/123 -v
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", users.SearchUser)

	// curl -X POST localhost:8000/users -v
	router.POST("/users", users.CreateUser)
}
