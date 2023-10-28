package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	runnerModel "github.com/kaikeventura/cat-runner/src/runner/model"
	"github.com/kaikeventura/cat-runner/src/storage/model"
	strategyModel "github.com/kaikeventura/cat-runner/src/strategy/model"
)

type StorageService struct{}

func ConstructStorageService() StorageService {
	return StorageService{}
}

func (StorageService) CreateStrategyTestFile(strategyTestName string) error {
	strategy := model.StrategyFile{
		StrategyTestName:     strategyTestName,
		CreatedAt:            time.Now(),
		HttpRequestRunners:   []runnerModel.HttpRunner{},
		EnvironmentVariables: []strategyModel.EnvironmentVariable{},
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
		return err
	}

	return nil
}

func (StorageService) FindAllStrategyTests() ([]model.StrategyFile, error) {
	directoryPath := getDirectoryPath()
	strategies := []model.StrategyFile{}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error when trying open directory:", err)
		return []model.StrategyFile{}, err
	}

	for _, file := range files {
		readFile, err := os.ReadFile(filepath.Join(directoryPath, file.Name()))
		if err != nil {
			fmt.Println("Error when trying read file:", err)
			return nil, err
		}

		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			var strategy = model.StrategyFile{}
			if err := json.Unmarshal(readFile, &strategy); err != nil {
				fmt.Println("Json parser error:", err)
				return nil, err
			}

			strategies = append(strategies, strategy)
		}
	}

	return strategies, nil
}

func (StorageService) FindStrategyByName(strategyTestName string) (*model.StrategyFile, error) {
	directoryPath := getDirectoryPath()
	filePath := filepath.Join(directoryPath, strategyTestName+".json")

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error when trying read file:", err)
		return nil, err
	}

	var strategyFile = model.StrategyFile{}

	if err := json.Unmarshal(file, &strategyFile); err != nil {
		fmt.Println("Json parser error:", err)
		return nil, err
	}

	return &strategyFile, nil
}

func (StorageService) UpdateStrategyTestFile(strategyTestName string, updatedStrategyFile model.StrategyFile) error {
	directoryPath := getDirectoryPath()
	filePath := filepath.Join(directoryPath, strategyTestName+".json")

	_, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error when trying read file:", err)
		return err
	}

	jsonData, err := json.Marshal(updatedStrategyFile)
	if err != nil {
		fmt.Println("JSON parser error:", err)
		return err
	}

	err = os.WriteFile(filePath, []byte(jsonData), 0644)
	if err != nil {
		fmt.Println("Error when trying write file with updated data:", err)
		return err
	}

	if strategyTestName != updatedStrategyFile.StrategyTestName {
		err = os.Rename(filePath, updatedStrategyFile.StrategyTestName)
		if err != nil {
			fmt.Println("Error when trying rename file:", err)
			return err
		}
	}

	return nil
}

func getDirectoryPath() string {
	return filepath.Join("src/storage/files")
}
