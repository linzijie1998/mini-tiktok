package model

type MongoComment struct {
	AuthorId   int64  `bson:"author_id"`
	VideoId    int64  `bson:"video_id"`
	Content    string `bson:"content"`
	CreateTime string `bson:"create_time"`
}
