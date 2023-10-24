package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/app/controller"
)

func RouterConfiguration(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	router.Use(cors.New(config))

	v1 := router.Group("v1")
	{
		runner := v1.Group("runner")
		{
			runner.POST("/http", controller.RunHttp)
		}
	}

	return router
}
