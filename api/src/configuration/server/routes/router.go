package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	runnerController "github.com/kaikeventura/cat-runner/src/runner/controller"
	strategyController "github.com/kaikeventura/cat-runner/src/strategy/controller"
)

func RouterConfiguration(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{
		"X-CSRF-Token",
		"X-Requested-With",
		"Accept", "Accept-Version",
		"Content-Length",
		"Content-MD5",
		"Content-Type",
		"Date",
		"X-Api-Version",
		"Authorization",
	}

	router.Use(cors.New(config))

	v1 := router.Group("v1")
	{
		runner := v1.Group("runner")
		{
			runner.POST("http", runnerController.RunHttp)
		}
		strategy := v1.Group("strategy")
		{
			strategy.POST("", strategyController.CreateStrategy)
			strategy.GET("", strategyController.GetAllStrategies)
			strategy.GET(":strategyName", strategyController.GetStrategyByName)
			strategy.POST(":strategyName/http", strategyController.AddHttpRunner)
			strategy.POST(":strategyName/env", strategyController.AddEnvironmentVariable)
			strategy.PUT(":strategyName/env", strategyController.UpdateEnvironmentVariable)
		}
	}

	return router
}
