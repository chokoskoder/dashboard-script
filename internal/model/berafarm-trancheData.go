package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerafarmTrancheData struct {
	ID                                            primitive.ObjectID `bson:"_id,omitempty"`
	TrancheVaultAddress                           *string            `bson:"trancheVaultAddress"`
	TransactionHash                               *string            `bson:"transactionHash"`
	Timestamp                                     *string            `bson:"timestamp"`
	TotalActualYieldVaultInvestment               *string            `bson:"totalActualYieldVaultInvestment"`
	TotalActualRewardVaultInvestment              *string            `bson:"totalActualRewardVaultInvestment"`
	TotalActualYieldVaultInvestmentAfterSlippage  *string            `bson:"totalActualYieldVaultInvestmentAfterSlippage"`
	TotalActualRewardVaultInvestmentAfterSlippage *string            `bson:"totalActualRewardVaultInvestmentAfterSlippage"`
	CurrentYieldVaultInvestment                   *string            `bson:"currentYieldVaultInvestment"`
	CurrentRewardVaultInvestment                  *string            `bson:"currentRewardVaultInvestment"`
	CurrentStrategyBalance                        *string            `bson:"currentStrategyBalance"`
	YieldTokenBalance                             *string            `bson:"yieldTokenBalance"`
	RewardTokenBalance                            *string            `bson:"rewardTokenBalance"`
	CreatedAt                                     time.Time          `bson:"createdAt"`
	UpdatedAt                                     time.Time          `bson:"updatedAt"`
}
