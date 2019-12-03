package infra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type MongoDb struct {
	client *mongo.Client
}

func (m *MongoDb) Collection(name string) *mongo.Collection {
	return m.client.Database("starwars").Collection(name)
}

var once sync.Once
func NewMongoDb() *MongoDb {
	var mongoDb *MongoDb

	once.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

		mongoDb = &MongoDb{client}
	})


	return mongoDb
}
