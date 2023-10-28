package service

import (
	"fmt"
	"time"

	runnerModel "github.com/kaikeventura/cat-runner/src/runner/model"
	storageModel "github.com/kaikeventura/cat-runner/src/storage/model"
	"github.com/kaikeventura/cat-runner/src/storage/service"
	"github.com/kaikeventura/cat-runner/src/strategy/model"
)

type StrategyService struct {
	storageService service.StorageService
}

func ConstructStrategyService(storageService service.StorageService) StrategyService {
	return StrategyService{storageService}
}

func (service StrategyService) CreateStrategy(strategy model.Strategy) error {
	err := service.storageService.CreateStrategyTestFile(strategy.Name)
	if err != nil {
		return err
	}

	return nil
}

func (service StrategyService) GetAllStrategies() ([]model.Strategy, error) {
	strategies, err := service.storageService.FindAllStrategyTests()
	if err != nil {
		return []model.Strategy{}, err
	}

	strategiesResponse := []model.Strategy{}

	for _, strategy := range strategies {
		var modifiedAt string
		if strategy.ModifiedAt != nil {
			modifiedAt = strategy.ModifiedAt.UTC().String()
		} else {
			modifiedAt = ""
		}

		strategiesResponse = append(
			strategiesResponse,
			model.Strategy{
				Name:       strategy.StrategyTestName,
				CreatedAt:  strategy.CreatedAt.UTC().String(),
				ModifiedAt: &modifiedAt,
			},
		)
	}

	return strategiesResponse, nil
}

func (service StrategyService) GetStrategyByName(strategyTestName string) (*storageModel.StrategyFile, error) {
	strategy, err := service.storageService.FindStrategyByName(strategyTestName)

	if err != nil {
		return nil, err
	}

	return strategy, nil
}

func (service StrategyService) AddHttpRunner(strategyTestName string, httpRunner runnerModel.HttpRunner) (*storageModel.StrategyFile, error) {
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

func (service StrategyService) AddEnvironmentVariable(strategyTestName string, environmentVariable model.EnvironmentVariable) (*storageModel.StrategyFile, error) {
	strategy, err := service.storageService.FindStrategyByName(strategyTestName)

	if err != nil {
		return nil, err
	}

	modifiedAt := time.Now()
	strategy.ModifiedAt = &modifiedAt

	for _, element := range strategy.EnvironmentVariables {
		if element.Name == environmentVariable.Name {
			return nil, fmt.Errorf("variable %s already exists", environmentVariable.Name)
		}
	}

	strategy.EnvironmentVariables = append(strategy.EnvironmentVariables, environmentVariable)

	service.storageService.UpdateStrategyTestFile(strategyTestName, *strategy)

	return strategy, nil
}

func (service StrategyService) UpdateEnvironmentVariable(strategyTestName string, environmentVariable model.EnvironmentVariable) (*storageModel.StrategyFile, error) {
	strategy, err := service.storageService.FindStrategyByName(strategyTestName)

	if err != nil {
		return nil, err
	}

	modifiedAt := time.Now()
	strategy.ModifiedAt = &modifiedAt

	for index, element := range strategy.EnvironmentVariables {
		if element.Name == environmentVariable.Name {
			strategy.EnvironmentVariables[index] = environmentVariable
			service.storageService.UpdateStrategyTestFile(strategyTestName, *strategy)
			return strategy, nil
		}
	}

	return nil, fmt.Errorf("variable %s not found", environmentVariable.Name)
}
