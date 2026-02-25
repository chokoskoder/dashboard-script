package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BerafarmReward struct {
	ID 						primitive.ObjectID 		`bson:"_id,omitempty"`

	// Pointers allow these to be strictly 'null' in the DB if missing.
	// We use 'omitempty' so they aren't stored as null, saving space.
	TransactionHash 			*string 			`bson:"transactionHash,omitempty"`

	// Address validation: We assume these are EVM addresses (0x...)
	//apparently we should add another layer of security here : a regex string checker
	TrancheVaultAddress  		*string 			`bson:"trancheVaultAddress,omitempty" validate:"omitempty,len=42"`
	StrategyAddress      		*string 			`bson:"strategyAddress,omitempty" validate:"omitempty,len=42"`
	RewardTokenAddress   		*string 			`bson:"rewardTokenAddress,omitempty" validate:"omitempty,len=42"`
	CurrencyTokenAddress 		string  			`bson:"currencyTokenAddress"` 

	//best to use decimal128 as floats cause calculation errors , which will cause a huge difference in calculations
	StrategyBalance      		primitive.Decimal128 `bson:"strategyBalance,omitempty"`
	RewardTokenAmount    		primitive.Decimal128 `bson:"rewardTokenAmount,omitempty"`
	RewardTrancheBalance		primitive.Decimal128 `bson:"rewardTrancheBalance,omitempty"`
	YieldTrancheBalance			primitive.Decimal128 `bson:"yieldTrancheBalance,omitempty"`
	RewardTokenPrice           	primitive.Decimal128 `bson:"rewardTokenPrice,omitempty"`
	RewardTokenPriceInCurrency 	primitive.Decimal128 `bson:"rewardTokenPriceInCurrency,omitempty"`
	CurrencyTokenPrice         	primitive.Decimal128 `bson:"currencyTokenPrice,omitempty"`
	RewardEarnedInCurrency     	primitive.Decimal128 `bson:"rewardEarnedInCurrency,omitempty"`

	CreatedAt 					time.Time 			`bson:"createdAt"`
	UpdatedAt 					time.Time 			`bson:"updatedAt"`
}