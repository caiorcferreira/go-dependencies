package infra

import (
	"context"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PlanetRepository struct {
	db *MongoDb
}

func (r PlanetRepository) GetAll(ctx context.Context) ([]Planet, error) {
	collection := r.db.Collection("planets")
	timeout, _ := context.WithTimeout(ctx, 20*time.Millisecond)

	cursor, err := collection.Find(timeout, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []Planet

	for cursor.Next(ctx) {
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
