package mongodb

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PublishMeta struct {
	PublishList []int64 `bson:"publish_list"`
}

func AddPublishInfo(ctx context.Context, uid, vid int64) error {
	publishCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("publish")
	filter := bson.M{"uid": uid}
	update := bson.M{
		"$addToSet": bson.M{
			"publish_list": bson.M{
				"$each": []int64{vid},
			},
		},
	}
	_, err := publishCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func GetPublishInfo(ctx context.Context, uid int64) ([]int64, error) {
	publishCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("publish")
	filter := bson.M{"uid": uid}
	var result PublishMeta
	if err := publishCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result.PublishList, nil
}
