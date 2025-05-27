package databases

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient interface {
	SetDB(name string) error
	Collection(name string) *mongo.Collection
	DC() error
}

type mongoDBClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDBConnection(dbURI string) (MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &mongoDBClient{client: client}, nil
}

func (c *mongoDBClient) SetDB(name string) error {
	c.db = c.client.Database(name)
	return nil
}

func (c *mongoDBClient) Collection(name string) *mongo.Collection {
	return c.db.Collection(name)
}

func (c *mongoDBClient) DC() error {
	ctx := context.TODO()
	err := c.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
