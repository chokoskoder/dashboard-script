package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseSchema struct {
	Apr            primitive.Decimal128 `bson:"apr"`
	PerformanceFee primitive.Decimal128 `bson:"performance_fee"`
}

type RewardsSchema struct {
	Multiplier     primitive.Decimal128 `bson:"multiplier"`
	PerformanceFee primitive.Decimal128 `bson:"performance_fee"`
}

type BaseVaultDetails struct {
	TrancheDisplayName  *string  `bson:"trancheDisplayName,omitempty"`
	StrategyInformation *string  `bson:"strategyInformation,omitempty"`
	VaultDescription    *string  `bson:"vaultDescription,omitempty"`
	VaultKeyPointers    []string `bson:"vaultKeyPointers,omitempty"`
}

type InvestmentToken struct {
	TokenAddress  *string `bson:"tokenAddress,omitempty"`
	TokenName     *string `bson:"tokenName,omitempty"`
	TokenSymbol   *string `bson:"tokenSymbol,omitempty"`
	TokenDecimals int32   `bson:"tokenDecimals"`
}

type EarnTokenDetails struct {
	Token *string `bson:"token,omitempty"`
	Value *string `bson:"value,omitempty"`
}

type VaultBerafarmTranche struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	TrancheVaultAddress    string             `bson:"trancheVaultAddress" validate:"omitempty,len=42"`
	StrategyManagerAddress string             `bson:"strategyManagerAddress" validate:"omitempty,len=42"`
	VaultManagerConfig     string             `bson:"vaultManagerConfig"`
	FeeModule              string             `bson:"feeModule"`
	AccessController       string             `bson:"accessController"`
	StrategyInformation    string             `bson:"strategyInformation"`
	VaultName              string             `bson:"vaultName"`
	YieldVaultDetails      BaseVaultDetails   `bson:"yieldVaultDetails"`
	RewardVaultDetails     BaseVaultDetails   `bson:"rewardVaultDetails"`
	VaultStatus            string             `bson:"vaultStatus"`
	YieldTokenName         string             `bson:"yieldTokenName"`
	YieldTokenSymbol       string             `bson:"yieldTokenSymbol"`
	YieldTokenAddress      string             `bson:"yieldTokenAddress" validate:"omitempty,len=42"`
	RewardTokenName        string             `bson:"rewardTokenName"`
	RewardTokenSymbol      string             `bson:"rewardTokenSymbol"`
	RewardTokenAddress     string             `bson:"rewardTokenAddress" validate:"omitempty,len=42"`
	IsActive               bool               `bson:"isActive"`
	BaseTokenAddress       string             `bson:"baseTokenAddress" validate:"omitempty,len=42"`
	BaseTokenSymbol        string             `bson:"baseTokenSymbol"`
	InvestmentTokens       []InvestmentToken  `bson:"investmentTokens,omitempty"`
	EarnTokenDetails       []EarnTokenDetails `bson:"earnTokenDetails,omitempty"`
	PointsFocused          string             `bson:"pointsFocused"`
	Fee                    string             `bson:"fee"`
	RewardTokenArray       []string           `bson:"rewardTokenArray,omitempty"`
	Tags                   []string           `bson:"tags,omitempty"`
	IsInTestingMode        bool               `bson:"isInTestingMode"`
	VaultSlippage          string             `bson:"vaultSlippage"`
	VaultType              string             `bson:"vaultType"`
	PointsArray            []string           `bson:"pointsArray,omitempty"`
	DataEmitterAddress     string             `bson:"dataEmitterAddress" validate:"omitempty,len=42"`
	CreatedAt              time.Time          `bson:"createdAt"`
	UpdatedAt              time.Time          `bson:"updatedAt"`
}