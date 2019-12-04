package infra

import (
	"context"
	. "github.com/caiorcferreira/swapi/internals/swapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PlanetRepository struct {
	db *MongoDb
}

type planetDto struct {
	Id primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Climate string `bson:"climate"`
	Terrain string `bson:"terrain"`
	Population string `bson:"population"`
}

func (r PlanetRepository) GetAll(ctx context.Context) ([]Planet, error) {
	collection := r.db.Collection("planets")
	//timeout, _ := context.WithTimeout(ctx, 50*time.Millisecond)

	cursor, err := collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []Planet

	for cursor.Next(ctx) {
		var item planetDto
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		p := Planet{
			Id: item.Id.Hex(),
			Name: item.Name,
			Terrain: item.Terrain,
			Climate: item.Climate,
			Population: item.Population,
		}

		results = append(results, p)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r PlanetRepository) Save(ctx context.Context, planet Planet) (Planet, error) {
	collection := r.db.Collection("planets")
	timeout, _ := context.WithTimeout(ctx, 20*time.Millisecond)

	result, err := collection.InsertOne(timeout, planet)
	if err != nil {
		return Planet{}, nil
	}

	objectId := result.InsertedID.(primitive.ObjectID).Hex()

	planet.Id = objectId

	return planet, nil
}

func NewPlanetRepository() PlanetRepository {
	db := NewMongoDb()

	return PlanetRepository{db}
}
