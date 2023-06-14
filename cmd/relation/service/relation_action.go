package service

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/relation/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/constant"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.ActionRequest) error {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	if claims.Id == req.ToUserId {
		return errno.ParamErr
	}
	if req.ActionType == constant.RelationActionDo {
		if err = db.CreateRelationInfos(s.ctx, []*model.Relation{
			{
				UserId:       claims.Id,
				FollowUserId: req.ToUserId,
			},
		}); err != nil {
			return err
		}
		return nil
	} else if req.ActionType == constant.RelationActionCancel {
		if err = db.DeleteRelationInfo(s.ctx, claims.Id, req.ToUserId); err != nil {
			return err
		}
		return nil
	} else {
		return errno.ParamErr
	}
}
