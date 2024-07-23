package utils

import (
	"go-nf/config"
	"log"
	"net"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type KafkaHandler struct {
	Conn *kafka.Conn
}

// Create Connection
func KafkaConn(cfg *config.KafkaConnCfg) *KafkaHandler {
	conn, err := kafka.Dial("tcp", cfg.Url)
	if err != nil {
		panic(err.Error())
	}
	// defer conn.Close()
	return &KafkaHandler{Conn: conn}
}

// Create Topic
func CreateTopic(conn *kafka.Conn) {
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	// Check topic if already exists or not
	var topicConfigs []kafka.TopicConfig
	for _, topic := range config.KafkaTopics {
		if !isTopicAlreadyExists(conn, topic) {
			topicConfigs = append(topicConfigs, kafka.TopicConfig{
				Topic:             topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			})
		}
	}

	if err := controllerConn.CreateTopics(topicConfigs...); err != nil {
		panic(err.Error())
	}
}

// Close Connection
func CloseConnection(conn *kafka.Conn) {
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// Validate Kafka Topic
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

func ListTopic(conn *kafka.Conn) []string {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	topics := []string{}
	for i := 0; i < len(partitions); i++ {
		topics = append(topics, partitions[i].Topic)
	}
	return topics
}

func (k *KafkaHandler) GetListTopic(c *fiber.Ctx) error {
	topics := ListTopic(k.Conn)
	return c.JSON(topics)
}

func (k *KafkaHandler) CreateTopics(c *fiber.Ctx) error {
	CreateTopic(k.Conn)
	return c.SendString("Created topics")
}

func (k *KafkaHandler) DeleteTopic(c *fiber.Ctx) error {
	k.Conn.DeleteTopics(ListTopic(k.Conn)...)
	return c.SendString("Deleted all topics")
}
