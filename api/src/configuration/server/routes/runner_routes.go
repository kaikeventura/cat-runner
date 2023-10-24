package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/runner/controller"
)

func runnerRoutes(baseGroup *gin.RouterGroup) {
	runner := baseGroup.Group("runner")
	{
		runner.POST("/http", controller.RunHttp)
	}
}
