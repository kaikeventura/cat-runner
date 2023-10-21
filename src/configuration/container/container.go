package container

import (
	"github.com/kaikeventura/cat-runner/src/app/client"
	"github.com/kaikeventura/cat-runner/src/app/controller"
	"github.com/kaikeventura/cat-runner/src/app/service"
)

func BuildConstainer() {
	buildRunner()
}

func buildRunner() {
	httpClient := client.ConstructHttpClient()
	runnerService := service.ConstructRunnerService(httpClient)
	controller.ConstructRunnerController(runnerService)
}
