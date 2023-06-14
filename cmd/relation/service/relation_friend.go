package service

import (
	"context"
	"errors"

	"github.com/linzijie1998/mini-tiktok/cmd/relation/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/pkg/constant"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"gorm.io/gorm"
)

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) FriendList(req *relation.FriendListRequest) ([]*relation.FriendUser, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	relationInfos, err := db.QueryFollowInfos(s.ctx, req.UserId, "follow_user_id")
	if err != nil {
		return nil, err
	}
	userList := make([]*relation.FriendUser, 0)
	for _, relationInfo := range relationInfos {
		err = db.QueryFollowInfo(s.ctx, relationInfo.FollowUserId, req.UserId, "id")
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		} else {
			userInfo, err := db.QueryFirstUserInfoByID(s.ctx, relationInfo.FollowUserId, "id, nickname, avatar")
			if err != nil {
				return nil, err
			}
			if len(userInfo.Avatar) == 0 {
				userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
			}
			message := "暂无聊天消息"
			msgType := constant.MessageTypeReceived
			latestSendMsgInfo, err := db.QueryLatestMessage(s.ctx, req.UserId, relationInfo.FollowUserId, "content, created_at")
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}
			} else {
				message = latestSendMsgInfo.Content
				msgType = constant.MessageTypeSend
			}
			latestReceivedMsgInfo, err := db.QueryLatestMessage(s.ctx, relationInfo.FollowUserId, req.UserId, "content, created_at")
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}
			} else {
				if latestReceivedMsgInfo.CreatedAt.After(latestSendMsgInfo.CreatedAt) {
					message = latestReceivedMsgInfo.Content
					msgType = constant.MessageTypeReceived
				}
			}
			userList = append(userList, &relation.FriendUser{
				Id:      userInfo.Id,
				Name:    userInfo.Nickname,
				Avatar:  &userInfo.Avatar,
				Message: &message,
				MsgType: int64(msgType),
			})
		}
	}
	return userList, nil
}
