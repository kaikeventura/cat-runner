package service

import (
	"time"

	runnerModel "github.com/kaikeventura/cat-runner/src/runner/model"
	"github.com/kaikeventura/cat-runner/src/storage/model"
	"github.com/kaikeventura/cat-runner/src/storage/service"
	strategyModel "github.com/kaikeventura/cat-runner/src/strategy/model"
)

type StrategyService struct {
	storageService service.StorageService
}

func ConstructStrategyService(storageService service.StorageService) StrategyService {
	return StrategyService{storageService}
}

func (service StrategyService) CreateStrategy(strategy strategyModel.Strategy) error {
	err := service.storageService.CreateStrategyTestFile(strategy.Name)
	if err != nil {
		return err
	}

	return nil
}

func (service StrategyService) GetAllStrategies() ([]string, error) {
	strategies, err := service.storageService.FindAllStrategyTests()
	if err != nil {
		return []string{}, err
	}

	return strategies, nil
}

func (service StrategyService) AddHttpRunner(strategyTestName string, httpRunner runnerModel.HttpRunner) (*model.StrategyFile, error) {
	strategy, err := service.storageService.FindStrategyByName(strategyTestName)

	if err != nil {
		return nil, err
	}

	if httpRunner.RequestName == "" {
		httpRunner.RequestName = "Http Request"
	}

	modifiedAt := time.Now()
	strategy.ModifiedAt = &modifiedAt

	strategy.HttpRequestRunners = append(strategy.HttpRequestRunners, httpRunner)

	service.storageService.UpdateStrategyTestFile(strategyTestName, *strategy)

	return strategy, nil
}
