package mongodb

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
)

type RelationMeta struct {
	FollowList   []int64 `bson:"follow_list"`
	FollowerList []int64 `bson:"follower_list"`
}

// AddRelationInfo 添加关注信息，通过mongo事务更新两个Documents：关注者的关注列表以及被关注者的粉丝列表
func AddRelationInfo(ctx context.Context, followerUid, followedUid int64) error {
	// 使用事务进行数据更新 https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#WithSession
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	session, err := global.MongoClient.StartSession(options.Session().SetDefaultReadConcern(readconcern.Majority()))
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	return mongo.WithSession(
		ctx,
		session,
		func(sessionContext mongo.SessionContext) error {
			opt := options.Update().SetUpsert(true)
			// 更新关注者的`follow_list`字段信息
			filter := bson.M{"uid": followerUid}
			update := bson.M{
				"$addToSet": bson.M{
					"follow_list": bson.M{
						"$each": []int64{followedUid},
					},
				},
			}
			_, err = relationCollection.UpdateOne(ctx, filter, update, opt)
			if err != nil {
				return err
			}
			// 更新被关注者`follower_list`的字段信息
			filter = bson.M{"uid": followedUid}
			update = bson.M{
				"$addToSet": bson.M{
					"follower_list": bson.M{
						"$each": []int64{followerUid},
					},
				},
			}
			_, err = relationCollection.UpdateOne(ctx, filter, update, opt)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

// GetFollowInfo 获取关注信息
func GetFollowInfo(ctx context.Context, uid, toUid int64) (bool, error) {
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	filter := bson.D{{"follow_list", bson.A{toUid}}}
	projection := bson.D{{"uid", uid}}
	err := relationCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteRelationInfo(ctx context.Context, followerUid, followedUid int64) error {
	// 使用事务进行数据更新 https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#WithSession
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	session, err := global.MongoClient.StartSession(options.Session().SetDefaultReadConcern(readconcern.Majority()))
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	return mongo.WithSession(
		ctx,
		session,
		func(sessionContext mongo.SessionContext) error {
			opt := options.Update().SetUpsert(true)
			// 更新关注者的`follow_list`字段信息
			filter := bson.M{"uid": followerUid}
			update := bson.D{{"$pull", bson.D{{"follow_list", followedUid}}}}
			_, err = relationCollection.UpdateOne(ctx, filter, update, opt)
			if err != nil {
				return err
			}
			// 更新被关注者`follower_list`的字段信息
			filter = bson.M{"uid": followedUid}
			update = bson.D{{"$pull", bson.D{{"follower_list", followerUid}}}}
			_, err = relationCollection.UpdateOne(ctx, filter, update, opt)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func GetFollowList(ctx context.Context, uid int64) ([]int64, error) {
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	filter := bson.M{"uid": uid}
	var result RelationMeta
	if err := relationCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result.FollowList, nil
}

func GetFollowerList(ctx context.Context, uid int64) ([]int64, error) {
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	filter := bson.M{"uid": uid}
	var result RelationMeta
	if err := relationCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result.FollowerList, nil
}

func GetFriendList(ctx context.Context, uid int64) ([]int64, error) {
	relationCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("relation")
	filter := bson.M{"uid": uid}
	var result RelationMeta
	if err := relationCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	if len(result.FollowList) > len(result.FollowerList) {
		result.FollowList, result.FollowerList = result.FollowerList, result.FollowList
	}
	m := map[int64]struct{}{}
	friendList := make([]int64, 0)
	for _, v := range result.FollowList {
		m[v] = struct{}{}
	}
	for _, v := range result.FollowerList {
		if _, ok := m[v]; ok {
			friendList = append(friendList, v)
		}
	}
	return friendList, nil
}
