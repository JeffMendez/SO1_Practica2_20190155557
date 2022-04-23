package main

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	//"log"
	//"net/http"	
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


func apiConnection(data string){
	responseBody := bytes.NewBuffer([]byte(data))
	fmt.Println("Consumio:", responseBody)

	/*resp, err := http.Post("http://34.133.28.136:3200/", "application/json", responseBody)

	if err != nil {
		log.Println("An Error Occured %v", err)
	}
	
	defer resp.Body.Close()*/
}
