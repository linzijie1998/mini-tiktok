// Code generated by Kitex v0.6.1. DO NOT EDIT.
package messageservice

import (
	server "github.com/cloudwego/kitex/server"
	message "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler message.MessageService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
