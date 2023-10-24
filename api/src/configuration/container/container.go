package container

import (
	"github.com/kaikeventura/cat-runner/src/runner/client"
	"github.com/kaikeventura/cat-runner/src/runner/controller"
	"github.com/kaikeventura/cat-runner/src/runner/service"
)

func BuildConstainer() {
	buildRunner()
}

func buildRunner() {
	httpClient := client.ConstructHttpClient()
	runnerService := service.ConstructRunnerService(httpClient)
	controller.ConstructRunnerController(runnerService)
}
