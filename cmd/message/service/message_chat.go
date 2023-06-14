package service

import (
	"context"
	"sort"

	"github.com/linzijie1998/mini-tiktok/cmd/message/dal/db"
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
	messages, err := db.QueryMessageByUserIDAndToUserIDWithLimit(s.ctx, req.ToUserId, claims.Id, req.PreMsgTime, "id, from_user_id, to_user_id, content, create_date")
	if err != nil {
		return nil, err
	}
	if req.PreMsgTime == 0 {
		// 查询发送给对方的消息
		sendMessages, err := db.QueryMessageByUserIDAndToUserIDWithLimit(s.ctx, claims.Id, req.ToUserId, req.PreMsgTime, "id, from_user_id, to_user_id, content, create_date")
		if err != nil {
			return nil, err
		}
		messages = append(sendMessages, messages...)
	}
	sort.Slice(messages, func(i, j int) bool {
		return (messages[i].CreatedAt).Before(messages[j].CreatedAt)
	})
	messageList := make([]*message.Message, len(messages))
	for i, msg := range messageList {
		messageList[i] = &message.Message{
			Id:         msg.Id,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Content,
			CreateDate: msg.CreateDate,
		}
	}
	return messageList, nil
}
