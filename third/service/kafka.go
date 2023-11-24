package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/segmentio/kafka-go"
)

func ListenToKafka() {
	topic := "micr_test_topic"
	brokers := []string{"localhost:29092"}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer reader.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		cancel()
	}()

	// Start consuming messages
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer stopped")
			return
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received message: key = %s, value = %s\n", string(msg.Key), string(msg.Value))
		}
	}
}
