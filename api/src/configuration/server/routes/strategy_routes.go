package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/strategy/controller"
)

func strategyRoutes(baseGroup *gin.RouterGroup) {
	strategy := baseGroup.Group("strategy")
	{
		strategy.POST("/", controller.CreateStrategy)
		strategy.GET("/", controller.GetAllStrategies)
	}
}
