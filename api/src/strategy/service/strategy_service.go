package service

import (
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

func (service StrategyService) GetAllStrategies() ([]string, error) {
	strategies, err := service.storageService.FindAllStrategyTests()
	if err != nil {
		return []string{}, err
	}

	return strategies, nil
}
