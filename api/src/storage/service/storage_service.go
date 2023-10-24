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

func CreateStrategyTestFile(StrategyTestName string) error {
	strategy := model.StrategyFile{
		StrategyTestName: StrategyTestName,
		CreatedAt:        time.Now(),
		RequestRunners:   &[]string{},
	}

	directoryPath := getDirectoryPath()
	filePath := filepath.Join(directoryPath, StrategyTestName+".json")
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

func FindAllStrategyTests() []string {
	directoryPath := getDirectoryPath()
	fileNames := []string{}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Error when trying open directory:", err)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			fileName := strings.TrimSuffix(file.Name(), ".json")
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames
}

func getDirectoryPath() string {
	return filepath.Join("src/storage/files")
}
