package mongodb

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/message/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddMessageInfo(ctx context.Context, message *model.MongoMessage) error {
	messageCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("message")
	_, err := messageCollection.InsertOne(ctx, message)
	return err
}

func GetAllMessages(ctx context.Context, uid1, uid2 int64) ([]*model.MongoMessage, error) {
	messageCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("message")
	filter := bson.M{
		"$or": []bson.M{
			{"sender": uid1, "receiver": uid2},
			{"sender": uid2, "receiver": uid1},
		},
	}
	cursor, err := messageCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var messages []*model.MongoMessage
	for cursor.Next(ctx) {
		var message model.MongoMessage
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func GetReceiveMessageWithLimit(ctx context.Context, sender, receiver int64, limit int64) ([]*model.MongoMessage, error) {
	messageCollection := global.MongoClient.Database(global.Configs.MongoDB.Database).Collection("message")
	filter := bson.M{
		"receiver": receiver,
		"sender":   sender,
		"timestamp": bson.M{
			"$gt": limit,
		},
	}
	cursor, err := messageCollection.Find(ctx, filter, options.Find().SetSort(bson.M{"timestamp": -1}))
	if err != nil {
		return nil, err
	}
	var messages []*model.MongoMessage
	for cursor.Next(ctx) {
		var message model.MongoMessage
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
