package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeeklyChange struct {
	Value     string `bson:"value"`
	Direction string `bson:"direction"`
}

type ReturnMetric struct {
	Value        string       `bson:"value"`
	Label        string       `bson:"label"`
	Badge        *string      `bson:"badge,omitempty"`
	Tooltip      *string      `bson:"tooltip,omitempty"`
	WeeklyChange WeeklyChange `bson:"weeklyChange"`
}

type CurrentYield struct {
	Apy       string `bson:"apy"`
	Apr       string `bson:"apr"`
	RewardApy string `bson:"rewardApy"`
	Label     string `bson:"label"`
}

type Tvl struct {
	Value primitive.Decimal128 `bson:"value"`
	Label string               `bson:"label"`
}

type Metrics struct {
	NetReturn      ReturnMetric `bson:"netReturn"`
	AbsoluteReturn ReturnMetric `bson:"absoluteReturn"`
	CurrentYield   CurrentYield `bson:"currentYield"`
	Tvl            Tvl          `bson:"tvl"`
}

type PerformancePoint struct {
	Day          string `bson:"day"`
	Date         string `bson:"date"`
	DailyGainPct string `bson:"dailyGainPct"`
	AvgGainPct   string `bson:"avgGainPct"`
	DailyAPY     string `bson:"dailyAPY"`
	AvgAPY       string `bson:"avgAPY"`
}

type Summary struct {
	TotalGainPct string `bson:"totalGainPct"`
	AvgDailyPct  string `bson:"avgDailyPct"`
	AvgAPY       string `bson:"avgAPY"`
}

type StrategySnapshot struct {
	Name       string               `bson:"name"`
	Allocation primitive.Decimal128 `bson:"allocation"`
	Status     string               `bson:"status"`
}

type TrancheTier struct {
	Allocation primitive.Decimal128 `bson:"allocation"`
	Apy        interface{}          `bson:"apy"`
	Type       string               `bson:"type"`
}

type Tranches struct {
	Senior TrancheTier `bson:"senior"`
	Junior TrancheTier `bson:"junior"`
}

type DashboardSnapshot struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	VaultName           string               `bson:"vaultName"`
	TrancheVaultAddress string               `bson:"trancheVaultAddress" validate:"required,len=42"`
	DisplayName         string               `bson:"displayName"`
	Description         string               `bson:"description"`
	Metrics             Metrics              `bson:"metrics"`
	YieldPerformance    []PerformancePoint   `bson:"yieldPerformance"`
	RewardPerformance   []PerformancePoint   `bson:"rewardPerformance"`
	YieldSummary        Summary              `bson:"yieldSummary"`
	RewardSummary       Summary              `bson:"rewardSummary"`
	Strategies          []StrategySnapshot   `bson:"strategies"`
	Tranches            Tranches             `bson:"tranches"`
	LastComputedAt      time.Time            `bson:"lastComputedAt"`
	CreatedAt           time.Time            `bson:"createdAt"`
	UpdatedAt           time.Time            `bson:"updatedAt"`
}
