package repositories

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	opt "go.mongodb.org/mongo-driver/mongo/options"
	"main.go/shared"
)

var (
	err *error
	ctx = context.TODO()
)

func ConnectDatabase() (*mongo.Database, error) {
	clientOptions := opt.Client().ApplyURI(shared.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error initializing database: %v", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		slog.Error("Ping error: %v", err)
	}
	slog.Info("Successfully connected to MongoDB!")
	return client.Database("ecTeste"), nil
}
