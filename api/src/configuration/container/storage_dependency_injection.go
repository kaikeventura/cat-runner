package container

import "github.com/kaikeventura/cat-runner/src/storage/service"

func buildStorageService() service.StorageService {
	return service.ConstructStorageService()
}
