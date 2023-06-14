// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/linzijie1998/mini-tiktok/cmd/api/global"
	"github.com/linzijie1998/mini-tiktok/cmd/api/initialize"
	"github.com/linzijie1998/mini-tiktok/cmd/api/initialize/rpc"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func LoadConfigsAndInit() {
	var err error
	if global.Viper, err = initialize.Viper("/home/nahida/devgo/src/mini-tiktok/cmd/api/config.yaml"); err != nil {
		panic(err)
	}
	if global.UserServiceClient, err = rpc.InitUserRPC(); err != nil {
		panic(err)
	}
	if global.PublishServiceClient, err = rpc.InitPublishRPC(); err != nil {
		panic(err)
	}
	if global.FeedServiceClient, err = rpc.InitFeedRPC(); err != nil {
		panic(err)
	}
	if global.FavoriteServiceClient, err = rpc.InitFavoriteRPC(); err != nil {
		panic(err)
	}
	if global.CommentServiceClient, err = rpc.InitCommentRPC(); err != nil {
		panic(err)
	}
	if global.RelationServiceClient, err = rpc.InitRelationRPC(); err != nil {
		panic(err)
	}
	if global.MessageServiceClient, err = rpc.InitMessageRPC(); err != nil {
		panic(err)
	}
}

func main() {
	LoadConfigsAndInit()

	h := server.Default(
		server.WithStreamBody(true),
		server.WithAltTransport(standard.NewTransporter),
		server.WithHostPorts(global.Configs.Hertz.Addr()),
	)

	h.Static("/videos", "/opt/tiktok")
	h.Static("/covers", "/opt/tiktok")

	register(h)
	h.Spin()
}
