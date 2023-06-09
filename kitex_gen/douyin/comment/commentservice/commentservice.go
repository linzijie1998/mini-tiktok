// Code generated by Kitex v0.5.2. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	comment "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/comment"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"comment_action": kitex.NewMethodInfo(commentActionHandler, newCommentServiceCommentActionArgs, newCommentServiceCommentActionResult, false),
		"comment_list":   kitex.NewMethodInfo(commentListHandler, newCommentServiceCommentListArgs, newCommentServiceCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentActionArgs)
	realResult := result.(*comment.CommentServiceCommentActionResult)
	success, err := handler.(comment.CommentService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentActionArgs() interface{} {
	return comment.NewCommentServiceCommentActionArgs()
}

func newCommentServiceCommentActionResult() interface{} {
	return comment.NewCommentServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentListArgs)
	realResult := result.(*comment.CommentServiceCommentListResult)
	success, err := handler.(comment.CommentService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentListArgs() interface{} {
	return comment.NewCommentServiceCommentListArgs()
}

func newCommentServiceCommentListResult() interface{} {
	return comment.NewCommentServiceCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *comment.ActionRequest) (r *comment.ActionResponse, err error) {
	var _args comment.CommentServiceCommentActionArgs
	_args.Req = req
	var _result comment.CommentServiceCommentActionResult
	if err = p.c.Call(ctx, "comment_action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *comment.ListRequest) (r *comment.ListResponse, err error) {
	var _args comment.CommentServiceCommentListArgs
	_args.Req = req
	var _result comment.CommentServiceCommentListResult
	if err = p.c.Call(ctx, "comment_list", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
