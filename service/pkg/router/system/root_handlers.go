package system

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/home-network-api/core/pkg/logger"
	"github.com/jgfranco17/home-network-api/service/pkg/data"
	"github.com/jgfranco17/home-network-api/service/pkg/env"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Algorithm API page!",
	})
}

func AboutHandler(c *gin.Context) {
	// Send the parsed JSON data as a response
	c.JSON(http.StatusOK, data.AboutInfo{
		Name:        "Algorithm API",
		Author:      "Joaquin Franco",
		Repository:  "https://github.com/jgfranco17/home-network-api",
		Version:     "0.0.1",
		Environment: env.GetApplicationEnv(),
		License:     "MIT",
		Languages:   []string{"Go"},
		Algorithms: map[string][]string{
			"array": {
				"MaxSubArray",
				"TwoSum",
			},
			"sequence": {
				"LongestCommonSubsequence",
				"Fibonacci",
			},
			"palindrome": {
				"palindrome",
			},
		},
	})
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, data.HealthStatus{
		Timestamp: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		Status:    "healthy",
	})
}

func NotFoundHandler(c *gin.Context) {
	log := logger.FromContext(c)
	log.Errorf("Non-existent endpoint accessed: %s", c.Request.URL.Path)
	c.JSON(http.StatusNotFound, newMissingEndpoint(c.Request.URL.Path))
}

func newMissingEndpoint(endpoint string) data.BasicErrorInfo {
	return data.BasicErrorInfo{
		StatusCode: http.StatusNotFound,
		Endpoint:   endpoint,
		Message:    "Endpoint does not exist",
	}
}
