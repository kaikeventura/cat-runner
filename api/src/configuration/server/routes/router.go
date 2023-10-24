package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	runnerController "github.com/kaikeventura/cat-runner/src/runner/controller"
	strategyController "github.com/kaikeventura/cat-runner/src/strategy/controller"
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
			runner.POST("/http", runnerController.RunHttp)
		}
		strategy := v1.Group("strategy")
		{
			strategy.POST("/", strategyController.CreateStrategy)
			strategy.GET("/", strategyController.GetAllStrategies)
			strategy.POST("/:strategyName/http", strategyController.AddHttpRunner)
		}
	}

	return router
}
