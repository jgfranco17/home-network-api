package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler() func(c *gin.Context) error {
	return func(c *gin.Context) error {
		username := c.Param("user")
		c.JSON(http.StatusOK, gin.H{
			"user":    username,
			"message": "Hello, world!",
		})
		return nil
	}
}
