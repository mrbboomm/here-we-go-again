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
		Brokers:     []string{brokerAddress},
		MaxBytes:    10e6, // 10MB
		Partition:   0,
		GroupTopics: topics,
		GroupID:     "my-group",
	}

	r := kafka.NewReader(readerConfig)
	// Consume messages from the topics
	for {
		msg, err := r.ReadMessage(c.Ctx)
		if err != nil {
			log.Printf("Error while reading message: %v", err)
			break
		}
		fmt.Printf("message from topic %s at offset %d: %s = %s\n", msg.Topic, msg.Offset, string(msg.Key), string(msg.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	consumer := &ConsumerHandler{BrokerAddress: os.Getenv("KAFKA_HOST"), Topics: config.KafkaTopics, Ctx: context.Background()}
	consumer.InitializeReader()
}
