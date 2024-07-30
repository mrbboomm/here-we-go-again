package repositories

import (
	"context"
	"go-nf/domains"
	"go-nf/entities"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type countryRepo struct {
	db *mongo.Client
}

func NewCountryRepo(db *mongo.Client) domains.CountryRepo {
	return &countryRepo{db}
}

func (c *countryRepo) Create(data entities.CountryEntity) error {
	collection := c.db.Database("testdb2").Collection("country")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}
	return nil
}
