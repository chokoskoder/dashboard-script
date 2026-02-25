package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerafarmApr struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	TrancheVaultAddress *string            `bson:"trancheVaultAddress"`
	Date                *string            `bson:"date"`
	CurrentRate         string             `bson:"currentRate"`
	InitialRate         string             `bson:"initialRate"`
	AbsoluteReturn      string             `bson:"absoluteReturn"`
	Apr                 string             `bson:"apr"`
	RewardReturn        string             `bson:"rewardReturn"`
	RewardApr           string             `bson:"rewardApr"`
	DaysSinceCreation   int                `bson:"daysSinceCreation"`
	RewardMetadata      interface{}        `bson:"rewardMetadata"`
	Timestamp           int64              `bson:"timestamp"`
	CreatedAt           time.Time          `bson:"createdAt"`
	UpdatedAt           time.Time          `bson:"updatedAt"`
}
