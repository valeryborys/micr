package service

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"fmt"
)

var brokerAddress string
var topic string

func SendMessage() {
	brokerAddress = "localhost:29092"
	topic = "micr_test_topic"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	defer writer.Close()

	message := kafka.Message{
		Key:   []byte("key"),
		Value: []byte("Hello, micr!"), 
	}

	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	fmt.Println("Message sent successfully to Kafka!")
}
