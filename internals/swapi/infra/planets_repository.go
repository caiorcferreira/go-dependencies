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

type planetDescription struct {
	Name string `bson:"name"`
	Climate string `bson:"climate"`
	Terrain string `bson:"terrain"`
	Population string `bson:"population"`
	Films int `bson:"films,omitempty"`
}

type planetDto struct {
	Id primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Climate string `bson:"climate"`
	Terrain string `bson:"terrain"`
	Population string `bson:"population"`
	Films int `bson:"films,omitempty"`
}

func (r PlanetRepository) GetAll(ctx context.Context) ([]Planet, error) {
	collection := r.db.Collection("planets")
	timeout, _ := context.WithTimeout(ctx, 100*time.Millisecond)

	cursor, err := collection.Find(timeout, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []Planet

	for cursor.Next(ctx) {
		var dto planetDto
		if err := cursor.Decode(&dto); err != nil {
			return nil, err
		}

		p := Planet{
			Id:         dto.Id.Hex(),
			Name:       dto.Name,
			Terrain:    dto.Terrain,
			Climate:    dto.Climate,
			Population: dto.Population,
			Films: dto.Films,
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
	timeout, _ := context.WithTimeout(ctx, 250*time.Millisecond)

	newPlanet := planetDescription{
		Name:       planet.Name,
		Population: planet.Population,
		Climate:    planet.Climate,
		Terrain:    planet.Terrain,
		Films: planet.Films,
	}

	result, err := collection.InsertOne(timeout, newPlanet)
	if err != nil {
		return Planet{}, err
	}

	objectId := result.InsertedID.(primitive.ObjectID).Hex()

	planet.Id = objectId

	return planet, nil
}

func NewPlanetRepository() PlanetRepository {
	db := NewMongoDb()

	return PlanetRepository{db}
}
