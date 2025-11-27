package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	URI        string
	Database   string
	Collection string
	client     *mongo.Client
	coll       *mongo.Collection
}

func NewMongo(uri, database, colleciotn string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	c := client.Database(database).Collection(colleciotn)
	return &MongoDB{
		URI:        uri,
		Database:   database,
		Collection: colleciotn,
		client:     client,
		coll:       c,
	}, nil
}

func (m *MongoDB) Name() string { return "mongo" }

func (m *MongoDB) WriteTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		_, err := m.coll.InsertOne(ctx, map[string]interface{}{
			"index": i,
			"value": fmt.Sprintf("value-%d", i),
			"ts":    time.Now().UnixNano(),
		})
		if err != nil {
			return 0, err
		}
	}
	return time.Since(start), nil
}

func (m *MongoDB) ReadTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		_ = m.coll.FindOne(ctx, map[string]interface{}{"index": i}).Err()
	}
	return time.Since(start), nil
}
