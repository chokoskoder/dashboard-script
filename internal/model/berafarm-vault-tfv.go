package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerafarmVaultTfv struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	TrancheVaultAddress   string             `bson:"trancheVaultAddress" validate:"required,len=42"`
	YieldVaultInvestment  string             `bson:"yieldVaultInvestment"`
	RewardVaultInvestment string             `bson:"rewardVaultInvestment"`
	CurrentTfv            string             `bson:"currentTfv"`
	Timestamp             string             `bson:"timestamp"`
}