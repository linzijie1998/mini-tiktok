package pack

import (
	"errors"

	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildActionResp(err error) *message.ActionResponse {
	if err == nil {
		return actionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return actionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return actionResp(s)
}

func actionResp(err errno.ErrNo) *message.ActionResponse {
	return &message.ActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildChatResp(messageList []*message.Message, err error) *message.ChatResponse {
	if err == nil {
		return chatResp(messageList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return chatResp(nil, e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return chatResp(nil, s)
}

func chatResp(messageList []*message.Message, err errno.ErrNo) *message.ChatResponse {
	return &message.ChatResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, MessageList: messageList}
}
