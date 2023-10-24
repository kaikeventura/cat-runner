package container

import (
	"github.com/kaikeventura/cat-runner/src/strategy/controller"
	"github.com/kaikeventura/cat-runner/src/strategy/service"
)

func buildStrategyController() {
	controller.ConstructStrategyService(buildStrategyService())
}

func buildStrategyService() service.StrategyService {
	return service.ConstructStrategyService(buildStorageService())
}
