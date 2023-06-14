package main

import (
	"context"
	"fmt"

	"github.com/linzijie1998/mini-tiktok/cmd/message/pack"
	"github.com/linzijie1998/mini-tiktok/cmd/message/service"
	message "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.ChatRequest) (*message.ChatResponse, error) {
	if len(req.Token) == 0 || req.ToUserId == 0 {
		return nil, errno.ParamErr
	}
	messageList, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		return pack.BuildChatResp(nil, err), nil
	}
	return pack.BuildChatResp(messageList, nil), nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.ActionRequest) (*message.ActionResponse, error) {
	fmt.Printf("%#v\n", req)
	if len(req.Token) == 0 || req.ToUserId == 0 {
		return nil, errno.ParamErr
	}
	if err := service.NewMessageActionService(ctx).MessageAction(req); err != nil {
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(nil), nil
}
