package model

type MongoMessage struct {
	Sender    int64  `bson:"sender"`
	Receiver  int64  `bson:"receiver"`
	Content   string `bson:"content"`
	Timestamp int64  `bson:"timestamp"`
}
