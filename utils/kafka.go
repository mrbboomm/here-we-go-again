package utils

import (
	"context"
	"go-nf/config"
	"log"

	"github.com/segmentio/kafka-go"
)

// Create Connection
func KafkaConn(cfg *config.KafkaConnCfg) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Url, cfg.Topic, 0)
	if err != nil {
		panic(err.Error())
	}

	// Check topic if already exists or not
	if !isTopicAlreadyExists(conn, cfg.Topic) {
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

	return conn
}

// Close Connection
func CloseConnection(conn *kafka.Conn) {
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func isTopicAlreadyExists(conn *kafka.Conn, topic string) bool {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partitions {
		if p.Topic == topic {
			return true
		}
	}
	return false
}
