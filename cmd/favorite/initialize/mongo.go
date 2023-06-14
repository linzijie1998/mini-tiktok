package initialize

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDB() (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(global.Configs.MongoDB.Url()))
	if err != nil {
		return nil, err
	}
	return client.Database(global.Configs.MongoDB.Database), nil
}
