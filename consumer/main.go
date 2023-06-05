package main

import (
	"awesomeProject1/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	config := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-events-v2",
		GroupID: "default-consumer",
	}
	reader := kafka.NewReader(config)
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		e := entity.Event{}
		err = json.Unmarshal(message.Value, &e)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("event:", e)
	}

}
