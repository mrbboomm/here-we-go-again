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

	"github.com/joho/godotenv"
)

func main() {
	//  Load env
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("NOT HAVE LOCAL ENV")
	}

	// Oracle Connection
	oracle.Connect()

	// Connection part
	cfg := config.KafkaConnCfg{
		Url:   os.Getenv("KAFKA_HOST"),
		Topic: "tier",
	}
	conn := utils.KafkaConn(&cfg)

	producer := &producer.ProducerHandler{Conn: conn}
	//Mock Data
	tiers := []tier.Tier{
		{
			Id:   2,
			Name: tier.Lang{En: "premium", Th: "พรีเมี่ยม"},
		},
	}
	// test Publish Event
	producer.PublishEvent(tiers)

	tier := &tier.Tier{Id: 1, Name: tier.Lang{En: "t", Th: "a"}}
	user := &user.User{Username: "hello", Password: "world", Tier: tier}
	fmt.Println("hello world")
	fmt.Println(user)
	fmt.Println(user.Tier)
}
