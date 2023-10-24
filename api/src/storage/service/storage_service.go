package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kaikeventura/cat-runner/src/storage/model"
)

type StorageService struct{}

func ConstructStorageService() StorageService {
	return StorageService{}
}

func (StorageService) CreateStrategyTestFile(strategyTestName string) error {
	strategy := model.StrategyFile{
		StrategyTestName: strategyTestName,
		CreatedAt:        time.Now(),
		RequestRunners:   &[]string{},
	}

	directoryPath := getDirectoryPath()
	filePath := filepath.Join(directoryPath, strategyTestName+".json")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error when trying create file", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(strategy); err != nil {
		fmt.Println("Error when trying write file:", err)
	}

	return nil
}

func (StorageService) FindAllStrategyTests() ([]string, error) {
	directoryPath := getDirectoryPath()
	fileNames := []string{}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error when trying open directory:", err)
		return []string{}, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			fileName := strings.TrimSuffix(file.Name(), ".json")
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames, nil
}

func getDirectoryPath() string {
	return filepath.Join("src/storage/files")
}
