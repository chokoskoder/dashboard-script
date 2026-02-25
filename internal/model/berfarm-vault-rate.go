package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerafarmVaultRate struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	TrancheVaultAddress string             `bson:"trancheVaultAddress" validate:"required,len=42"`
	YieldVaultRate      string             `bson:"yieldVaultRate"`
	RewardVaultRate     string             `bson:"rewardVaultRate"`
	YieldMultiplier     string             `bson:"yieldMultiplier"`
	RewardMultiplier    string             `bson:"rewardMultiplier"`
	Timestamp           string             `bson:"timestamp"`
}