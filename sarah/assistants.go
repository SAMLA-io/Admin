package sarah

import (
	"context"
	"log"
	"os"
	"samla-admin/types/mongodb"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[SARAH] WARNING: .env file not found, using system environment variables")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Pinged deployment. Successfully connected to MongoDB!")
}

func GetOrganizationAssistants(orgId string) ([]mongodb.Assistant, error) {
	coll := Client.Database(orgId).Collection(os.Getenv("MONGO_COLLECTION_ASSISTANTS"))

	cursor, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var assistants []mongodb.Assistant
	if err := cursor.All(context.Background(), &assistants); err != nil {
		log.Println(err)
		return nil, err
	}

	return assistants, nil
}
