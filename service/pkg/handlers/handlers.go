package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.Param("user")
		c.JSON(http.StatusOK, gin.H{
			"user":    username,
			"message": "Hello, world!",
		})
	}
}
