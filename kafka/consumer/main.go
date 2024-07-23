package main

import (
	"context"
	"fmt"
	"go-nf/config"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type ConsumerHandler struct {
	BrokerAddress string
	Topics        []string
	Ctx           context.Context
}

func (c *ConsumerHandler) InitializeReader() {
	// Kafka broker address
	brokerAddress := c.BrokerAddress

	// Topics to subscribe to
	topics := c.Topics

	// Kafka reader configuration
	readerConfig := kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  "my-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	// Create Kafka readers for each topic
	readers := make([]*kafka.Reader, len(topics))
	for i, topic := range topics {
		config := readerConfig
		config.Topic = topic
		readers[i] = kafka.NewReader(config)
	}

	defer c.CloseReader(readers)
	// Consume messages from the topics
	for {
		fmt.Printf("waiting message for %v reader \n", len(readers))
		for _, reader := range readers {
			msg, err := reader.ReadMessage(c.Ctx)
			if err != nil {
				log.Printf("Error while reading message: %v", err)
				continue
			}
			fmt.Printf("Received message from %s: %s\n", msg.Topic, string(msg.Value))
			fmt.Printf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))
		}
		// Add a small delay to avoid tight looping
		// time.Sleep(1 * time.Second)
	}
}

func (c *ConsumerHandler) CloseReader(readers []*kafka.Reader) {
	// Close the readers when done
	for _, reader := range readers {
		if err := reader.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}
}

func main() {
	consumer := &ConsumerHandler{BrokerAddress: os.Getenv("KAFKA_HOST"), Topics: config.KafkaTopics, Ctx: context.Background()}
	consumer.InitializeReader()
}
