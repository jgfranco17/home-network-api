package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/home-network-api/service/pkg/handlers"
	error_handling "github.com/jgfranco17/home-network-api/service/pkg/router/error_handling"
)

// Adds v0 routes to the router.
func SetRoutes(route *gin.Engine) {
	v0 := route.Group("/v0")
	{
		v0.GET("greet/:user", error_handling.WithErrorHandling(handlers.WelcomeHandler()))
	}
}
