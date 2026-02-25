package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Strategy struct {
	StrategyAddress           string               `bson:"strategyAddress"`
	StrategyName              *string              `bson:"strategyName,omitempty"`
	StrategyUnderlyingBalance primitive.Decimal128 `bson:"strategyUnderlyingBalance"`
	StrategyTokenBalance      primitive.Decimal128 `bson:"strategyTokenBalance"`
	StrategyRatio             primitive.Decimal128 `bson:"strategyRatio"`
}

type BerafarmStrategyBalance struct {
	ID                     primitive.ObjectID   `bson:"_id,omitempty"`
	TrancheVaultAddress    string               `bson:"trancheVaultAddress" validate:"required,len=42"`
	StrategyManagerAddress string               `bson:"strategyManagerAddress" validate:"required,len=42"`
	Strategies             []Strategy           `bson:"strategies,omitempty"`
	TotalBalance           primitive.Decimal128 `bson:"totalBalance"`
	CreatedAt              time.Time            `bson:"createdAt"`
	UpdatedAt              time.Time            `bson:"updatedAt"`
}