package main

import (
	"fmt"
	"go-nf/config"
	"go-nf/kafka/producer"
	oracle "go-nf/oracle/connection"
	"go-nf/tier"
	"go-nf/user"
	"go-nf/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//  Load env
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("NOT HAVE LOCAL ENV")
	}

	KAFKA_HOST := os.Getenv("KAFKA_HOST")
	// Connection part
	oracleCfg := config.LocalOracleConfig()
	oracle.Connect(&oracleCfg)

	cfg := config.KafkaConnCfg{
		Url:    KAFKA_HOST,
		Topics: config.KafkaTopics,
	}
	kafkaHandler := utils.KafkaConn(&cfg)

	// Check topics
	if topics := utils.ListTopic(kafkaHandler.Conn); len(topics) == 0 {
		utils.CreateTopic(kafkaHandler.Conn)
	}

	tier := &tier.Tier{Id: 1, Name: tier.Lang{En: "t", Th: "a"}}
	user := &user.User{Username: "hello", Password: "world", Tier: tier}
	fmt.Println("hello world")
	fmt.Println(user)
	fmt.Println(user.Tier)

	// Initialize Fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/kafka/list-topic", kafkaHandler.GetListTopic)
	app.Post("/kafka/topic", kafkaHandler.CreateTopics)
	app.Delete("/kafka/topic", kafkaHandler.DeleteTopic)
	app.Post("/kafka/producer", producer.SendMassage)

	app.Listen(":3000")
}
