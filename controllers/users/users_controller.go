package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	// take json body from request
	// // the below block of code can be replaced with
	// // c.ShouldBindJSON(&user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO handle error
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		// TODO return user creation error
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if err != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusCreated, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me plz")
}
