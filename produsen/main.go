package main

import (
	"awesomeProject1/entity"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-events-v2",
	})
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	msgs := make([]kafka.Message, 0, 100)
	for i := 0; i < 100; i++ {
		e := entity.Event{
			Name: "updated",
			Data: entity.User{
				ID:   uint(i + 1),
				Name: "john",
			},
		}

		data, err := json.Marshal(e)
		if err != nil {
			log.Fatalln(err)
		}

		message := kafka.Message{
			Value: data,
		}
		msgs = append(msgs, message)
	}

	err := writer.WriteMessages(context.Background(), msgs...)
	if err != nil {
		log.Println(err)
	}
}
