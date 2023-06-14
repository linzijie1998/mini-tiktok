package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/dal/cache"
	"os"
	"path/filepath"

	"github.com/linzijie1998/mini-tiktok/cmd/publish/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/publish"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/ffmpeg"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"gorm.io/gorm"
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
	// 2. 计算HASH值, 查找是否有相同的视频
	hash := fmt.Sprintf("%x", sha256.Sum256(req.Data))
	videoInfo, err := db.QueryVideoInfoByHash(s.ctx, hash, "id, video_path, cover_path")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有相同视频
		videoInfo = &model.Video{
			Hash:      hash,
			VideoPath: fmt.Sprintf("%s.mp4", hash), // 以hash值为文件名
			CoverPath: fmt.Sprintf("%s.jpg", hash), // 以hash值为文件名
			AuthorId:  claims.Id,
			Title:     req.Title,
		}
		videoPath := filepath.Join(global.Configs.Upload.VideoPath, videoInfo.VideoPath)
		coverPath := filepath.Join(global.Configs.Upload.CoverPath, videoInfo.CoverPath)
		// 保存文件
		file, err := os.Create(videoPath)
		if err != nil {
			return errno.ServiceErr.WithMessage(err.Error())
		}
		if _, err = file.Write(req.Data); err != nil {
			return errno.ServiceErr.WithMessage(err.Error())
		}
		defer file.Close()
		// 截取封面
		if err = ffmpeg.GetCover(videoPath, coverPath, "00:00:00"); err != nil {
			return errno.ServiceErr.WithMessage(err.Error())
		}
	} else {
		// 找到相同视频
		if err != nil {
			return err
		}
		videoInfo.DefaultModel = model.DefaultModel{}
		videoInfo.Hash = hash
		videoInfo.AuthorId = claims.Id
		videoInfo.Title = req.Title
	}
	// 3. 存储视频上传信息
	if err = db.AddPublishInfo(s.ctx, claims.Id, videoInfo); err != nil {
		return errno.ServiceErr.WithMessage(err.Error())
	}
	// 4. 添加缓存信息
	if err = cache.NewVideoInfos(s.ctx, []*model.Video{videoInfo},
		global.Configs.CacheExpire.ParseVideoBaseInfoExpireDuration()); err != nil {
		return err
	}
	if err = cache.NewVideoCounters(s.ctx, []*model.Video{videoInfo}); err != nil {
		return err
	}
	if err = cache.AddPublishInfo(s.ctx, claims.Id, videoInfo.Id); err != nil {
		return err
	}
	if err = cache.IncrWorkCount(s.ctx, claims.Id); err != nil {
		return err
	}
	return nil
}
