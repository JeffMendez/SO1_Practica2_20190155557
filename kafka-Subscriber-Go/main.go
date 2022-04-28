package main

import (
	"fmt"
	"strconv"
	"strings"
	"subscriber/db"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
    fmt.Println("201901557 - Subscriber iniciado")

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
			insertarDBs(string(msg.Value))
		} else {			
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}

func insertarDBs(data string){
	dataLog := strings.Split(data, "\t")
	
	if len(dataLog) == 3 {
		gameID, _ := strconv.Atoi(dataLog[0])
		players, _ := strconv.Atoi(dataLog[1])
		winner, _ := strconv.Atoi(dataLog[2])
		gamename := ""

		switch gameID {
		case 1:
			gamename = "Ruleta"
		case 2:
			gamename = "Dados"
		case 3:
			gamename = "Dardos"
		case 4:
			gamename = "CartaMayor"
		case 5:
			gamename = "SillasMusicales"
		}

		db.InsertarLogMongo(gameID, players, winner, gamename)
	}
}
