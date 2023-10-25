package model

import (
	"time"

	runnerModel "github.com/kaikeventura/cat-runner/src/runner/model"
	"github.com/kaikeventura/cat-runner/src/strategy/model"
)

type StrategyFile struct {
	StrategyTestName     string
	CreatedAt            time.Time
	ModifiedAt           *time.Time
	HttpRequestRunners   []runnerModel.HttpRunner
	EnvironmentVariables []model.EnvironmentVariable
}
