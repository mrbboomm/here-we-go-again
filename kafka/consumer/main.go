package main

import (
	"fmt"
	"go-nf/config"
	"go-nf/utils"

	"github.com/segmentio/kafka-go"
)

type ConsumerHandler struct {
	Conn *kafka.Conn
}

func (c *ConsumerHandler) SubscribeEvent() {
	for {
		message, err := c.Conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}
}

func main() {
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "tier",
	}
	conn := utils.KafkaConn(&cfg)

	consumer := &ConsumerHandler{Conn: conn}
	consumer.SubscribeEvent()
}
