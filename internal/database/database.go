package database

import "go.mongodb.org/mongo-driver/mongo"

//we will connect to the database here and use that conection instance in the repository folder to work with the db
//connect to mongoDB
//we setup the mongodb connection uri string in our config file , meaning we will call it here and work with it.

func SetupDBConnection(dbName ,dbUri string ) mongo.Connect {
	
}