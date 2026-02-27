package repository

import (
	"context"
	"errors"

	"github.com/chokoskoder/dashboard-script/internal/database"
	"github.com/chokoskoder/dashboard-script/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Define standard errors (optional but recommended for service layer checks)
var (
	ErrNotFound = errors.New("document not found")
)

// VaultRepository defines the contract for vault-related data access.
// Renamed from VaultInfo to VaultRepository (or VaultReader) to signify its role.
type VaultRepository interface {
	GetTVL(ctx context.Context, address string) (*model.BerafarmVaultTfv, error)
	GetStrategies(ctx context.Context, address string) (*model.BerafarmStrategyBalance, error)
	GetTrancheAllocation(ctx context.Context, address string) (*model.BerafarmTrancheData, error)
}

// mongoVaultRepository implementation of the VaultRepository interface.
// It holds references to all collections needed to fulfill the interface.
type mongoVaultRepository struct {
	tvlCollection      *mongo.Collection
	strategyCollection *mongo.Collection
	trancheCollection  *mongo.Collection
}

// GetStrategies implements VaultRepository.
}

// Config struct to avoid messy function signatures if params grow
type MongoConfig struct {
	DBName             string
	TVLCollection      string
	StrategyCollection string
	TrancheCollection  string
}

// NewMongoVaultRepository creates a single repository that handles vault data.
func NewMongoVaultRepository(client *mongo.Client, cfg MongoConfig) VaultRepository {
	return &mongoVaultRepository{
		tvlCollection:      database.CollectionConn(client, cfg.DBName, cfg.TVLCollection),
		strategyCollection: database.CollectionConn(client, cfg.DBName, cfg.StrategyCollection),
		trancheCollection:  database.CollectionConn(client, cfg.DBName, cfg.TrancheCollection),
	}
}

// Implementation Example: GetTVL
func (r *mongoVaultRepository) GetTVL(ctx context.Context, address string) (*model.BerafarmVaultTfv, error) {
	var result model.BerafarmVaultTfv

	// Example filter - assuming address is the _id or a field
	filter := map[string]string{"address": address}

	err := r.tvlCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound // Return a domain-specific error
		}
		return nil, err
	}

	return &result, nil
}
func (r *mongoVaultRepository) GetStrategies(ctx context.Context, address string) (*model.BerafarmStrategyBalance, error) {
	panic("unimplemented")
}

// GetTrancheAllocation implements VaultRepository.
func (r *mongoVaultRepository) GetTrancheAllocation(ctx context.Context, address string) (*model.BerafarmTrancheData, error) {
	panic("unimplemented")

// You would then implement GetStrategies and GetTrancheAllocation similarly...
