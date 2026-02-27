package repository

import (
	"context"
	"fmt"

	"github.com/chokoskoder/dashboard-script/internal/database"
	"github.com/chokoskoder/dashboard-script/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

//this is where we will get all the data from our apr collections.

type APRRepository interface {
	GetAPR(ctx context.Context) ([]model.BerafarmVaultRate , error)
}

type mongoAPRRepository struct {
	collection *mongo.Collection
}

func NewAPRRepository(client *mongo.Client , dbName , collectionName string) *mongoAPRRepository {

	collectionConn := database.CollectionConn(client , dbName , collectionName)
	return &mongoAPRRepository{
		collection: collectionConn,
	}
}

func(r *mongoAPRRepository) GetAPR(ctx context.Context) ([]model.BerafarmVaultRate , error){
	//we will need to get apr based on what exaclty ?
	// the logic we are implementing will need us to work with dates right ?? -> so how do we query this ?   
	cursor , err := r.collection.Find(ctx , bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	var APR []model.BerafarmVaultRate

	// cursor.All automatically iterates over all results and decodes them 
	// into the slice. It maps BSON fields -> Struct fields.
	if err := cursor.All(ctx, &APR); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	return APR, nil
	//this apr will be used for calculations ?? which apr are we calling and from where ??
}