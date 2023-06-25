package mongodb

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FavoriteMeta struct {
	FavoriteList []int64 `bson:"favorite_list"`
}

func AddFavoriteInfo(ctx context.Context, uid, vid int64) error {
	favoriteCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("favorite")
	filter := bson.M{"uid": uid}
	update := bson.M{
		"$addToSet": bson.M{
			"favorite_list": bson.M{
				"$each": []int64{vid},
			},
		},
	}
	_, err := favoriteCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func GetFavoriteList(ctx context.Context, uid int64) ([]int64, error) {
	favoriteCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("favorite")
	filter := bson.M{"uid": uid}
	var result FavoriteMeta
	if err := favoriteCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result.FavoriteList, nil
}

func GetFavoriteInfo(ctx context.Context, uid, toUid int64) (bool, error) {
	favoriteCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("favorite")
	filter := bson.D{{"favorite_list", bson.A{toUid}}}
	projection := bson.D{{"uid", uid}}
	err := favoriteCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteFavoriteInfo(ctx context.Context, uid, vid int64) error {
	// 使用事务进行数据更新 https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#WithSession
	var err error
	favoriteCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("favorite")
	opt := options.Update().SetUpsert(true)
	// 更新关注者的`follow_list`字段信息
	filter := bson.M{"uid": uid}
	update := bson.D{{"$pull", bson.D{{"favorite_list", vid}}}}
	_, err = favoriteCollection.UpdateOne(ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}
