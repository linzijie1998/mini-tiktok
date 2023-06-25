package initialize

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func Mongo() (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(global.Configs.MongoDB.Url()))
}
