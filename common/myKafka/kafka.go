package myKafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func CreateKafkaConnection() *kafka.Conn {
	// Replace with your Kafka broker addresses
	brokers := []string{"localhost:9092"}

	// Create a new Kafka connection
	conn, _ := kafka.DialContext(context.Background(), "tcp", brokers[0])

	return conn
}

func SendMessage(message string, topic string, conn *kafka.Conn) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Write the message to Kafka
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func ConsumeMessages(topic string, partition int, conn *kafka.Conn) {
	// Create a new reader for the specified partition
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		MaxWait:   5 * time.Second,
	})

	// Start an infinite loop to continuously consume messages
	for {
		// Read the next message from Kafka
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		// Process the received message
		processMessage(msg.Value)
	}
}

func processMessage(message []byte) {
	// TODO: Implement your message processing logic here
	fmt.Println("Received message:", string(message))
}
