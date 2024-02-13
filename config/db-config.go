package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Naveenchand06/go-gin-jwt-auth/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func GetDB() *mongo.Client {
	if db != nil {
		return db
	} else {
		// * Reading mongoDB URL from .env
		mongoUrl := os.Getenv(constants.MONGOURL)
		
		// * Creating a `context` with TimeOut
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// * The below line to connects to MongoDb and returns `*mongo.Client`
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
		if err != nil {
			log.Fatalln(err)
		}
		db = client
		log.Println("Connected to MongoDB")
		return db
	}
}