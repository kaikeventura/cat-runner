package model

import "time"

type StrategyFile struct {
	StrategyTestName string
	CreatedAt        time.Time
	ModifiedAt       *time.Time
	RequestRunners   *[]string
}
