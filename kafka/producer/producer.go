package producer

import (
	"go-nf/utils"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type ProducerHandler struct {
	Conn *kafka.Conn
}

// Method PublishEvent
func (p *ProducerHandler) PublishEvent(msgs ...any) {
	// Convert into kafka.Message{}
	// TODO move to utils
	messages := make([]kafka.Message, 0)
	for _, p := range msgs {
		messages = append(messages, kafka.Message{
			Value: utils.CompressToJsonBytes(&p),
		})
	}

	// Set timeout
	p.Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := p.Conn.WriteMessages(messages...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

// Method CloseConnection
func (p *ProducerHandler) CloseConnection() {
	if err := p.Conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
