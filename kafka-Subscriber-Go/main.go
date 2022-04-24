package main

import (
	"fmt"
	"strings"
	//"io/ioutil"
	//"log"
	//"net/http"
	"strconv"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"	
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
    fmt.Println("Subscriber iniciado")
	
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-bootstrap",
		"group.id":          "grupo201901557",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{"juegos", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))			
			apiConnection(string(msg.Value))
		} else {			
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}

func connect(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
					30 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){ 
	defer cancel()
	defer func() {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	}()
}

func ping(client *mongo.Client, ctx context.Context) error{
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
    collection := client.Database(dataBase).Collection(col) 
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}

func apiConnection(data string){
	dataLog := strings.Split(data, "\t")
	
	if len(dataLog) == 3 {
		gameID, _ := strconv.Atoi(dataLog[0])
		players, _ := strconv.Atoi(dataLog[1])
		winner, _ := strconv.Atoi(dataLog[2])

		client, ctx, cancel, err := connect("mongodb://mongoadmin:So1pass1S_2022@34.125.124.39:27017/LogsMongo?authSource=admin&ssl=false")
		if err != nil {
			panic(err)
		}
		defer close(client, ctx, cancel)

		var document interface{}
		document = bson.D{
			{"game_id", gameID},
			{"players", players},
			{"game_name", "-"},
			{"winner", winner},
			{"queue", "Kafka"},
		}

    	_, errAdd := insertOne(client, ctx, "LogsMongo", "logs", document)
		if errAdd != nil {
			panic(errAdd)
		}
	}
}
