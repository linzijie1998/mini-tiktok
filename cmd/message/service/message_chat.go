package service

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/message/dal/mongodb"
	"github.com/linzijie1998/mini-tiktok/cmd/message/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *message.ChatRequest) ([]*message.Message, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}

	if req.PreMsgTime == 0 {
		// 查询所有聊天记录
		messages, err := mongodb.GetAllMessages(s.ctx, claims.Id, req.ToUserId)
		if err != nil {
			return nil, err
		}
		messageList := make([]*message.Message, len(messages))
		for i, mongoMessage := range messages {
			messageList[i] = &message.Message{
				Id:         int64(i + 10),
				ToUserId:   mongoMessage.Receiver,
				FromUserId: mongoMessage.Sender,
				Content:    mongoMessage.Content,
				CreateTime: &mongoMessage.Timestamp,
			}
		}
		return messageList, nil
	}
	messages, err := mongodb.GetReceiveMessageWithLimit(s.ctx, req.ToUserId, claims.Id, req.PreMsgTime)
	if err != nil {
		return nil, err
	}
	messageList := make([]*message.Message, len(messages))
	for i, mongoMessage := range messages {
		messageList[i] = &message.Message{
			Id:         int64(i + 100),
			ToUserId:   mongoMessage.Receiver,
			FromUserId: mongoMessage.Sender,
			Content:    mongoMessage.Content,
			CreateTime: &mongoMessage.Timestamp,
		}
	}
	return messageList, nil
}
