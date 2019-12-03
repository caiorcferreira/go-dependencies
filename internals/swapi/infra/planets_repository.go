package infra

import (
	"context"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlanetRepository struct {
	db *MongoDb
}

func (r PlanetRepository) GetAll() ([]Planet, error) {
	c := r.db.Collection("planets")

	cursor, err := c.Find(context.TODO(), bson.D{}, options.Find())
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	var results []Planet

	for cursor.Next(context.TODO()) {
		var item Planet
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func NewPlanetRepository() PlanetRepository {
	db := NewMongoDb()

	return PlanetRepository{db}
}
