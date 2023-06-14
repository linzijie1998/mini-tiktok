package pack

import (
	"errors"

	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/comment"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildActionResp(commentInfo *comment.Comment, err error) *comment.ActionResponse {
	if err == nil {
		return actionResp(commentInfo, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return actionResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return actionResp(nil, s)
}

func actionResp(commentInfo *comment.Comment, err errno.ErrNo) *comment.ActionResponse {
	return &comment.ActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, Comment: commentInfo}
}

func BuildListResp(commentList []*comment.Comment, err error) *comment.ListResponse {
	if err == nil {
		return listResp(commentList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listResp(nil, s)
}

func listResp(commentList []*comment.Comment, err errno.ErrNo) *comment.ListResponse {
	return &comment.ListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, CommentList: commentList}
}
