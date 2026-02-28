package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WeeklyChange struct {
	Value     string `bson:"value" json:"value"`
	Direction string `bson:"direction" json:"direction"`
}

type MetricDetail struct {
	Value        string        `bson:"value" json:"value"`
	Label        string        `bson:"label" json:"label"`
	Badge        string        `bson:"badge,omitempty" json:"badge,omitempty"`
	Tooltip      string        `bson:"tooltip,omitempty" json:"tooltip,omitempty"`
	WeeklyChange *WeeklyChange `bson:"weeklyChange,omitempty" json:"weeklyChange,omitempty"`
}

type CurrentYieldMetric struct {
	Apy       string `bson:"apy" json:"apy"`
	Apr       string `bson:"apr" json:"apr"`
	RewardApy string `bson:"rewardApy" json:"rewardApy"`
	Label     string `bson:"label" json:"label"`
}

type TvlMetric struct {
	Value float64 `bson:"value" json:"value"`
	Label string  `bson:"label" json:"label"`
}

type VaultMetrics struct {
	NetReturn      MetricDetail       `bson:"netReturn" json:"netReturn"`
	AbsoluteReturn MetricDetail       `bson:"absoluteReturn" json:"absoluteReturn"`
	CurrentYield   CurrentYieldMetric `bson:"currentYield" json:"currentYield"`
	Tvl            TvlMetric          `bson:"tvl" json:"tvl"`
}

type DataPoint struct {
	Day          string `bson:"day" json:"day"`
	Date         string `bson:"date" json:"date"`
	DailyGainPct string `bson:"dailyGainPct" json:"dailyGainPct"`
	AvgGainPct   string `bson:"avgGainPct" json:"avgGainPct"`
	DailyAPY     string `bson:"dailyAPY" json:"dailyAPY"`
	AvgAPY       string `bson:"avgAPY" json:"avgAPY"`
}

type VaultSummary struct {
	TotalGainPct string `bson:"totalGainPct" json:"totalGainPct"`
	AvgDailyPct  string `bson:"avgDailyPct" json:"avgDailyPct"`
	AvgAPY       string `bson:"avgAPY" json:"avgAPY"`
}

type StrategyDashboard struct {
	ID         int     `bson:"id" json:"id"`
	Name       string  `bson:"name" json:"name"`
	Allocation float64 `bson:"allocation" json:"allocation"`
	Status     string  `bson:"status" json:"status"`
}

type TrancheTier struct {
	Allocation float64     `bson:"allocation" json:"allocation"`
	Apy        interface{} `bson:"apy" json:"apy"` // Schema.Types.Mixed in Mongoose
	Type       string      `bson:"type" json:"type"`
}

type TrancheData struct {
	Senior TrancheTier `bson:"senior" json:"senior"`
	Junior TrancheTier `bson:"junior" json:"junior"`
}

type DashboardSnapshot struct {
	ID                primitive.ObjectID 	`bson:"_id,omitempty" json:"-"`
	VaultID           string             	`bson:"id" json:"id"` // Maps to Mongoose "id"
	Name              string             	`bson:"name" json:"name"`
	Description       string             	`bson:"description,omitempty" json:"description,omitempty"`
	Metrics           VaultMetrics       	`bson:"metrics" json:"metrics"`
	YieldPerformance  []DataPoint        	`bson:"yieldPerformance" json:"yieldPerformance"`
	RewardPerformance []DataPoint        	`bson:"rewardPerformance" json:"rewardPerformance"`
	YieldSummary      VaultSummary       	`bson:"yieldSummary" json:"yieldSummary"`
	RewardSummary     VaultSummary       	`bson:"rewardSummary" json:"rewardSummary"`
	Strategies        []StrategyDashboard	`bson:"strategies" json:"strategies"`
	Tranches          TrancheData        	`bson:"tranches" json:"tranches"`
	CreatedAt         time.Time          	`bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          	`bson:"updatedAt" json:"updatedAt"`
}
