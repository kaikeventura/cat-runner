package model

import "time"

type StrategyTest struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt *time.Time
}
