package global

import (
	"github.com/linzijie1998/mini-tiktok/cmd/api/config"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/comment/commentservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/favorite/favoriteservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed/feedservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message/messageservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/publish/publishservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation/relationservice"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user/userservice"
	"github.com/spf13/viper"
)

var (
	Configs config.ServiceConfigs
	Viper   *viper.Viper

	UserServiceClient     *userservice.Client
	PublishServiceClient  *publishservice.Client
	FeedServiceClient     *feedservice.Client
	FavoriteServiceClient *favoriteservice.Client
	CommentServiceClient  *commentservice.Client
	MessageServiceClient  *messageservice.Client
	RelationServiceClient *relationservice.Client
)
