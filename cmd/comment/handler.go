package main

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/comment/pack"
	"github.com/linzijie1998/mini-tiktok/cmd/comment/service"
	comment "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/comment"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.ActionRequest) (*comment.ActionResponse, error) {
	if len(req.Token) == 0 || req.VideoId == 0 {
		return pack.BuildActionResp(nil, errno.ParamErr), nil
	}
	commentInfo, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		return pack.BuildActionResp(nil, err), nil
	}
	return pack.BuildActionResp(commentInfo, nil), nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.ListRequest) (*comment.ListResponse, error) {
	if req.VideoId == 0 {
		return pack.BuildListResp(nil, errno.ParamErr), nil
	}
	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		return pack.BuildListResp(nil, err), nil
	}
	return pack.BuildListResp(commentList, nil), nil
}
