package app

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	mapURLs()

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
