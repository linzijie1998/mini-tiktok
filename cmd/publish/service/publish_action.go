package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/dal/cache"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/dal/mongodb"
	"github.com/linzijie1998/mini-tiktok/pkg/path"
	"os"
	"path/filepath"

	"github.com/linzijie1998/mini-tiktok/cmd/publish/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/publish"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/ffmpeg"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *publish.ActionRequest) error {
	// 1. 解析Token
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	// 2. 生成信息视频
	newId := uuid.NewString()
	videoInfo := &model.Video{
		VideoPath: fmt.Sprintf("videos/%d/%s.mp4", claims.Id, newId),
		CoverPath: fmt.Sprintf("covers/%d/%s.jpg", claims.Id, newId),
		AuthorId:  claims.Id,
		Title:     req.Title,
	}
	videoPath := filepath.Join(global.Configs.FileAccess.UploadPath, videoInfo.VideoPath)
	coverPath := filepath.Join(global.Configs.FileAccess.UploadPath, videoInfo.CoverPath)
	// 3. 保存文件
	if err = path.MakeDirs(videoPath); err != nil {
		return err
	}
	if err = path.MakeDirs(coverPath); err != nil {
		return err
	}
	file, err := os.Create(videoPath)
	if err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	if _, err = file.Write(req.Data); err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	defer file.Close()
	// 4. 截取封面
	if err = ffmpeg.GetCover(videoPath, coverPath, "00:00:00"); err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	// 5. 存储视频上传信息
	if err = db.AddPublishInfo(s.ctx, claims.Id, videoInfo); err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	if err = mongodb.AddPublishInfo(s.ctx, claims.Id, videoInfo.Id); err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	// 6. 添加缓存信息
	if err = cache.PushVideoQueue(s.ctx, videoInfo.Id, 30); err != nil {
		return err
	}
	if err = cache.NewVideoInfos(s.ctx, []*model.Video{videoInfo},
		global.Configs.CacheExpire.ParseVideoBaseInfoExpireDuration()); err != nil {
		return err
	}
	if err = cache.NewVideoCounters(s.ctx, []*model.Video{videoInfo}); err != nil {
		return err
	}
	//if err = cache.AddPublishInfo(s.ctx, claims.Id, videoInfo.Id); err != nil {
	//	return err
	//}
	if err = cache.IncrWorkCount(s.ctx, claims.Id); err != nil {
		return err
	}
	if err = cache.DelPublishInfoNullKey(s.ctx, claims.Id); err != nil {
		return err
	}

	return nil
}
