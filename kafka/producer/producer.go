package producer

import (
	"context"
	"fmt"
	"go-nf/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type ProducerHandler struct {
	Writer *kafka.Writer
}

type EventPayload struct {
	Topic   string `json:"topic"`
	Message any    `json:"message"`
}

// Initialize Producer
func Initialize(topic string) *ProducerHandler {
	brokerAddress := os.Getenv("KAFKA_HOST")
	log.Printf("broker address %s with topic %s \n", brokerAddress, topic)

	// init Writer
	w := &kafka.Writer{
		Addr:     kafka.TCP([]string{brokerAddress}...),
		Balancer: &kafka.LeastBytes{},
	}
	return &ProducerHandler{Writer: w}
}

// Method PublishEvent
func (p *ProducerHandler) PublishEvent(payloads ...EventPayload) string {
	// Convert into kafka.Message{}
	// TODO move to utils
	messages := make([]kafka.Message, 0)
	for _, payload := range payloads {
		messages = append(messages, kafka.Message{
			Topic: payload.Topic,
			Value: utils.CompressToJsonBytes(&payload.Message),
		})
	}
	ctx := context.Background()
	err := p.Writer.WriteMessages(ctx, messages...)
	if err != nil {
		log.Printf("Failed to write message: %v", err)
	}

	if err := p.Writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	fmt.Printf("Successfully produced message: %v\n", payloads)
	return "Successfully produced message"
}

func SendMassage(c *fiber.Ctx) error {
	payload := new(EventPayload)
	c.BodyParser(payload)

	writer := Initialize(payload.Topic)
	return c.JSON(writer.PublishEvent([]EventPayload{*payload}...))
}
