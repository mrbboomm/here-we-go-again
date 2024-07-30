package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type CountryEntity struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Continent string `json:"continent" bson:"continent"`
}