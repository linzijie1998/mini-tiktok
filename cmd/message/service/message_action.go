package service

import (
	"context"
	"time"

	"github.com/linzijie1998/mini-tiktok/cmd/message/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/message/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/message"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/constant"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(req *message.ActionRequest) error {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	if req.ActionType != constant.MessageChatTypeSend {
		return errno.ParamErr
	}
	newMessage := &model.Message{
		ToUserId:   req.ToUserId,
		FromUserId: claims.Id,
		Content:    req.Content,
		CreateTime: time.Now().Format("2006/01/02 15:04:05"),
	}
	if err = db.CreateMessageInfos(s.ctx, []*model.Message{newMessage}); err != nil {
		return err
	}
	return nil
}
