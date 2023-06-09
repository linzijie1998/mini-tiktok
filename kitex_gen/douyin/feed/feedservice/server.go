// Code generated by Kitex v0.5.2. DO NOT EDIT.
package feedservice

import (
	server "github.com/cloudwego/kitex/server"
	feed "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler feed.FeedService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
