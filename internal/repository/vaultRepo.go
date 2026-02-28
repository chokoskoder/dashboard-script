package repository

import (
	"context"
	"fmt"

	"github.com/chokoskoder/dashboard-script/internal/database"
	"github.com/chokoskoder/dashboard-script/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type vaultRepository interface {
	GetActiveVaults(ctx context.Context) ([]model.VaultBerafarmTranche, error)
}

type mongoVaultRepository struct {
	collection *mongo.Collection
}

func NewVaultRepository(client *mongo.Client, dbName, collectionName string) *mongoVaultRepository {
	collectionConn := database.CollectionConn(client, dbName, collectionName)
	return &mongoVaultRepository{
		collection: collectionConn,
	}
}

func (r *mongoVaultRepository) GetActiveVaults(ctx context.Context) ([]model.VaultBerafarmTranche, error) {
	// Filter for vaults where isActive is true
	filter := bson.M{"isActive": true}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find active vaults: %w", err)
	}
	defer cursor.Close(ctx)

	var vaults []model.VaultBerafarmTranche

	// cursor.All automatically iterates over all results and decodes them
	// into the slice. It maps BSON fields -> Struct fields.
	if err := cursor.All(ctx, &vaults); err != nil {
		return nil, fmt.Errorf("failed to decode vaults: %w", err)
	}

	return vaults, nil
}
