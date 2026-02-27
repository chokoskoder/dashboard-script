package repository

import (
	"context"
	"fmt"
	"github.com/chokoskoder/dashboard-script/internal/database"
	"github.com/chokoskoder/dashboard-script/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type dashboardRepository interface {
	GetDashboardData(ctx context.Context) ([]model.DashboardSnapshot)
	SaveDashboardData(ctx context.Context) ([]model.DashboardSnapshot)
}

type mongoDashboardRepository struct {
	collection *mongo.Collection
}

func NewDashboardRepository(client *mongo.Client , dbName , collectionName string) *mongoDashboardRepository {

	collectionConn := database.CollectionConn(client , dbName , collectionName)
	return &mongoDashboardRepository{
		collection: collectionConn,
	}
}

func(r *mongoDashboardRepository) SaveDashboardData(ctx context.Context , data model.DashboardSnapshot) error {
	_, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("Error inserting data into DashboardSnapshot model : %w" , err)
		//use custom error handling here
	}
	return nil
}

func (r *mongoDashboardRepository) GetDashboardData(ctx context.Context) ([]model.DashboardSnapshot , error) {
	cursor , err := r.collection.Find(ctx , bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	// we close the cursor when the function finishes
	//this cursor is just like what we studied in mysql python in 11th
	defer cursor.Close(ctx)
	var snapshots []model.DashboardSnapshot

	// cursor.All automatically iterates over all results and decodes them 
	// into the slice. It maps BSON fields -> Struct fields.
	if err := cursor.All(ctx, &snapshots); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	return snapshots, nil
}