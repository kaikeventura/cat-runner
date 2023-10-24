package model

import (
	"time"

	"github.com/kaikeventura/cat-runner/src/runner/model"
)

type StrategyFile struct {
	StrategyTestName   string
	CreatedAt          time.Time
	ModifiedAt         *time.Time
	HttpRequestRunners []model.HttpRunner
}
