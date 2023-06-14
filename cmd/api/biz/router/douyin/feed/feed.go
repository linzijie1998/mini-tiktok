// Code generated by hertz generator. DO NOT EDIT.

package feed

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	feed "github.com/linzijie1998/mini-tiktok/cmd/api/biz/handler/douyin/feed"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), feed.Feed)...)
		}
	}
}
