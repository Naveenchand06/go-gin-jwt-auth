package database

import (
	"os"

	"github.com/Naveenchand06/go-gin-jwt-auth/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(os.Getenv(constants.DBNAME)).Collection(collectionName)
}