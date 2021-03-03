package users

import (
	"bookstore_oauth-go/oauth"
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

func getUserId(userIDParam string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("id should be a number")
	}
	return userID, nil
}

// Create func passes the context to services
func Create(c *gin.Context) {
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

	result, err := services.UserService.CreateUser(user)
	if err != nil {
		// TODO return user creation error
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get func passes the context to services
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
	}
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UserService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	if oauth.GetCallerId(c.Request) == user.ID {
		c.JSON(http.StatusCreated, user.Marshall(false))
		return
	}

	c.JSON(http.StatusCreated, user.Marshall(oauth.IsPublic(c.Request)))
}

// Update func passes the context to services
func Update(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	res, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, res.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete func passes the context to services
func Delete(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search func passes the context to services
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// Login func passes the context to services
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, user.Marshall(c.GetHeader("X-Public") == "true"))
}
