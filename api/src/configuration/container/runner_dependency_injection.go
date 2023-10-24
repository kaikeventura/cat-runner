package container

import (
	"github.com/kaikeventura/cat-runner/src/runner/client"
	"github.com/kaikeventura/cat-runner/src/runner/controller"
	"github.com/kaikeventura/cat-runner/src/runner/service"
)

func buildRunnerController() {
	controller.ConstructRunnerController(buildRunnerService())
}

func buildRunnerService() service.RunnerService {
	return service.ConstructRunnerService(buildHttpClient())
}

func buildHttpClient() client.HttpClient {
	return client.ConstructHttpClient()
}
