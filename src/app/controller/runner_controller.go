package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-runner/src/app/model"
)

func RunHttp(context *gin.Context) {
	var runner model.HttpRunner

	err := context.ShouldBindJSON(&runner)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	context.JSON(201, runner)
}
