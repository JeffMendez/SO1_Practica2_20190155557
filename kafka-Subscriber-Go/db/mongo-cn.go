package db

import (
	"time"
	"fmt"
	"context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"	
	"go.mongodb.org/mongo-driver/bson"
)

func ConnectMongo(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
					30 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func CloseMongo(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){ 
	defer cancel()
	defer func() {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	}()
}

func InsertarMongo(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
    collection := client.Database(dataBase).Collection(col) 
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}

func InsertarLogMongo(gameID int, players int, winner int, gamename string) {
	client, ctx, cancel, err := ConnectMongo("mongodb://mongoadmin:So1pass1S_2022@34.125.211.171:27017/LogsMongo?authSource=admin&ssl=false")
		
	if err != nil {
		panic(err)
	}
	defer CloseMongo(client, ctx, cancel)

	collection := client.Database("LogsMongo").Collection("logs") 
	totalDatos, _ := collection.CountDocuments(ctx, bson.D{})

	var document interface{}
	document = bson.D{
		{"request_number", totalDatos},
		{"game", gameID},
		{"gamename", gamename},
		{"winner", winner},
		{"players", players},
		{"worker", "Kafka"},
	}
	fmt.Println("Mongo", document)

	_, errAdd := InsertarMongo(client, ctx, "LogsMongo", "logs", document)
	if errAdd != nil {
		panic(errAdd)
	}
}