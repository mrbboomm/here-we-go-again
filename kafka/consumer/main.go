package main

import (
	"fmt"
	"log"

	"go-nf/config"
	"go-nf/utils"
)

func main() {
	cfg := config.KafkaConnCfg{
		Url:   "localhost:9092",
		Topic: "tier",
	}
	conn := utils.KafkaConn(cfg)

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
