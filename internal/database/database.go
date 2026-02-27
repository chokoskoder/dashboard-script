package database

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//we will connect to the database here and use that conection instance in the repository folder to work with the db
//connect to mongoDB
//we setup the mongodb connection uri string in our config file , meaning we will call it here and work with it.

type Database struct {
	client *mongo.Client
}

func SetupDBConnection(dbUri , environment string, ctx context.Context , logger *slog.Logger ) (*Database , error) {

	//need to add logging here 
	serverAPi := options.ServerAPI(options.ServerAPIVersion1)
	//ok this is weird , why do we need to pass this by address ??
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty: true,
		OmitEmpty: true,
	}//its a struct thats why
	opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPi).SetBSONOptions(bsonOpts)

	client , err := mongo.Connect(opts)
	if err != nil{
		//add logger here
		return nil , fmt.Errorf("error while connecting to db %w" , err)
	}

	//ping the db
	if ping_err := client.Ping(ctx , nil); err != nil{
		return nil,fmt.Errorf("erro while pinging the db: %w" , ping_err)
	}

	return &Database{
		client: client,
	} ,nil

}

//graceful shutdown function 
// we accept a context here so we can enforce a timeout on the shutdown logic
func(db *Database ) Close( ctx context.Context) error{
	if db.client == nil {
		return nil
	}
	return db.client.Disconnect(ctx)
}