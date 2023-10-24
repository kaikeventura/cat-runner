package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/strategy/model"
	"github.com/kaikeventura/cat-runner/src/strategy/service"
)

var strategyService service.StrategyService

func ConstructStrategyService(service service.StrategyService) {
	strategyService = service
}

func CreateStrategyTest(context *gin.Context) {
	var strategy model.Strategy

	err := context.ShouldBindJSON(&strategy)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	context.JSON(201, strategyService.CreateStrategyTest(strategy))
}
