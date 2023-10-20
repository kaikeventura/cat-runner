package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/app/controller"
)

func RouterConfiguration(router *gin.Engine) *gin.Engine {
	v1 := router.Group("v1")
	{
		runner := v1.Group("runner")
		{
			runner.POST("/http", controller.RunHttp)
		}
	}

	return router
}
