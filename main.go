package main

import (
	"context"
	"fmt"
	"go-nf/config"
	"go-nf/deliveries"
	"go-nf/kafka/producer"
	"go-nf/mongodb"
	repositories "go-nf/repositories/country"
	"go-nf/tier"
	usecases "go-nf/usecases/country"
	"go-nf/user"
	"go-nf/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//  Load env
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("NOT HAVE LOCAL ENV")
	}

	KAFKA_HOST := os.Getenv("KAFKA_HOST")
	// Connection part
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

	// connect mongo
	clientOptions := options.Client().ApplyURI("mongodb://mongouser:mongopass@localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	mongodb.SetClient(client)

	// user api (TEST)
	// mongodb.ConnectToMongo()
	app.Post("/create-user", mongodb.CreateUserLogin)
	app.Get("/user", mongodb.GetAllUserLogin)
	app.Get("/user/:username", mongodb.GetUserLoginByUsername)
	app.Get("/user-id/:id", mongodb.GetUserLoginById)
	app.Put("/update-user/:id", mongodb.UpdateUserLoginById)
	app.Delete("/delete-user/:id", mongodb.DeleteUserLoginById)

// country api (TEST)
	countryRepo := repositories.NewCountryRepo(client)
	countryUseCase := usecases.NewCountryUseCase(countryRepo)
	countryHandlers := deliveries.NewCountryHandler((countryUseCase))

	app.Post("/country", countryHandlers.CreateCountry)

	app.Listen(":3000")

}
