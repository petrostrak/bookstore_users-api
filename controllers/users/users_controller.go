package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me plz")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me plz")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me plz")
}
