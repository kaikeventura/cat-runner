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

func (service StrategyService) CreateStrategyTest(strategy model.Strategy) string {
	return ""
}
