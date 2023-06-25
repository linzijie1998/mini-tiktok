package cache

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"strconv"
	"time"
)

func buildVideoInfoMap(video *model.Video) map[string]interface{} {
	if video == nil {
		return nil
	}
	return map[string]interface{}{
		"id":         video.Id,
		"author_id":  video.AuthorId,
		"video_path": video.VideoPath,
		"cover_path": video.CoverPath,
		"title":      video.Title,
	}
}

func parseVideoInfo(videoInfoMap map[string]string) (*model.Video, error) {
	videoInfo := new(model.Video)
	if value, ok := videoInfoMap["id"]; ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		videoInfo.Id = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	if value, ok := videoInfoMap["author_id"]; !ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		videoInfo.AuthorId = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	var ok bool
	if videoInfo.VideoPath, ok = videoInfoMap["video_path"]; !ok {
		return nil, errors.New("missing field")
	}
	if videoInfo.CoverPath, ok = videoInfoMap["cover_path"]; !ok {
		return nil, errors.New("missing field")
	}
	if videoInfo.Title, ok = videoInfoMap["title"]; !ok {
		return nil, errors.New("missing field")
	}
	return videoInfo, nil
}

func buildVideoCounterMap(video *model.Video) map[string]interface{} {
	if video == nil {
		return nil
	}
	return map[string]interface{}{
		"id":             video.Id,
		"favorite_count": video.FavoriteCount,
		"comment_count":  video.CommentCount,
	}
}

func NewVideoInfos(ctx context.Context, videos []*model.Video, duration time.Duration) error {
	pipe := global.RedisClient.Pipeline()
	for _, video := range videos {
		key := getVideoInfoKey(video.Id)
		if _, err := pipe.Del(ctx, key).Result(); err != nil {
			return err
		}
		videoInfoMap := buildVideoInfoMap(video)
		if _, err := pipe.HMSet(ctx, key, videoInfoMap).Result(); err != nil {
			return err
		}
		if _, err := pipe.Expire(ctx, key, duration).Result(); err != nil {
			return err
		}
	}
	if _, err := pipe.Exec(ctx); err != nil {
		// 报错后进行一次额外尝试
		if _, err = pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func NewVideoCounters(ctx context.Context, videos []*model.Video) error {
	pipe := global.RedisClient.Pipeline()
	for _, video := range videos {
		key := getVideoCounterKey(video.Id)
		if _, err := pipe.Del(ctx, key).Result(); err != nil {
			return err
		}
		videoCounterMap := buildVideoCounterMap(video)
		if _, err := pipe.HMSet(ctx, key, videoCounterMap).Result(); err != nil {
			return err
		}
	}
	if _, err := pipe.Exec(ctx); err != nil {
		// 报错后进行一次额外尝试
		if _, err = pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func GetVideoInfo(ctx context.Context, vid int64) (*model.Video, error) {
	key := getVideoInfoKey(vid)
	videoInfoMap, err := global.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return parseVideoInfo(videoInfoMap)
}

func GetVideoCounter(ctx context.Context, vid int64) (*model.Video, error) {
	key := getVideoCounterKey(vid)
	videoCounterMap, err := global.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return parseVideoInfo(videoCounterMap)
}
