package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterConfiguration(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	router.Use(cors.New(config))

	v1 := router.Group("v1")
	{
		runnerRoutes(v1)
		strategyRoutes(v1)
	}

	return router
}
