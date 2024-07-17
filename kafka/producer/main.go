package main

import (
	"go-nf/config"
	"go-nf/tier"
	"go-nf/utils"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Connection part
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "tier",
	}
	conn := utils.KafkaConn(cfg)

	// Check topic if already exists or not
	if !utils.IsTopicAlreadyExists(conn, cfg.Topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             cfg.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	// Mock data
	data := func() []kafka.Message {
		// Name: tier.Lang{En: "standard", Th: "เริ่มต้น"},
		tiers := []tier.Tier{
			{
				Id:   2,
				Name: tier.Lang{En: "premium", Th: "พรีเมี่ยม"},
			},
		}
		log.Printf("Mock message data => %v", tiers)
		// Convert into kafka.Message{}
		messages := make([]kafka.Message, 0)
		for _, p := range tiers {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJsonBytes(&p),
			})
		}
		return messages
	}()

	// Set timeout
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// Close connection
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
