package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/home-network-api/service/pkg/env"
)

const (
	API_URL_LOCAL = "http://localhost:3000"
	API_URL_DEV   = "https://dev.home-network-api.com"
	API_URL_STAGE = "https://stage.home-network-api.com"
	API_URL_BETA  = "https://beta.home-network-api.com"
	API_URL_PROD  = "https://www.home-network-api.com"
)

func getAllowedOrigins(environment string) []string {
	// TODO: Once more stable, remove local -> dev/stage/prod. For now it is useful for quick testing.
	switch environment {
	case env.APPLICATION_ENV_DEV:
		return []string{API_URL_LOCAL, API_URL_DEV}
	case env.APPLICATION_ENV_STAGE:
		return []string{API_URL_LOCAL, API_URL_STAGE}
	case env.APPLICATION_ENV_PROD:
		return []string{API_URL_LOCAL, API_URL_BETA, API_URL_PROD}
	default:
		return []string{API_URL_LOCAL}
	}
}

func GetCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
	})
}
