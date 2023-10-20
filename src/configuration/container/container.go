package container

import (
	"github.com/kaikeventura/cat-runner/src/app/controller"
	"github.com/kaikeventura/cat-runner/src/app/service"
)

func BuildConstainer() {
	buildRunner()
}

func buildRunner() {
	runnerService := service.ConstructRunnerService()
	controller.ConstructRunnerController(runnerService)
}
