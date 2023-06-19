package cache

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
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

func parseVideoCounter(videoCounterMap map[string]string) (*model.Video, error) {
	videoInfo := new(model.Video)
	if value, ok := videoCounterMap["id"]; ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		videoInfo.Id = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	if value, ok := videoCounterMap["favorite_count"]; ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		videoInfo.FavoriteCount = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	if value, ok := videoCounterMap["comment_count"]; ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		videoInfo.CommentCount = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	return videoInfo, nil
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

func AddPublishInfo(ctx context.Context, uid, vid int64) error {
	key := getPublishKey(uid)
	_, err := global.RedisClient.SAdd(ctx, key, vid).Result()
	return err
}

func NewPublishInfo(ctx context.Context, uid int64, vidList []int64) error {
	key := getPublishKey(uid)
	pipe := global.RedisClient.Pipeline()
	if _, err := pipe.Del(ctx, key).Result(); err != nil {
		return err
	}
	for _, vid := range vidList {
		if _, err := pipe.SAdd(ctx, key, vid).Result(); err != nil {
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

func GetPublishInfo(ctx context.Context, uid int64) ([]int64, error) {
	var err error
	key := getPublishKey(uid)
	result, err := global.RedisClient.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	vidList := make([]int64, len(result))
	for i, val := range result {
		vidList[i], err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return vidList, nil
}

func GetFavoriteStatus(ctx context.Context, uid, vid int64) (bool, error) {
	key := getVideoFavoriteKey(uid)
	return global.RedisClient.SIsMember(ctx, key, vid).Result()
}

func GetPublishInfoNullKey(ctx context.Context, uid int64) error {
	key := getPublishInfoNullKey(uid)
	return getNullKey(ctx, key)
}

func DelPublishInfoNullKey(ctx context.Context, uid int64) error {
	key := getPublishInfoNullKey(uid)
	return delNullKey(ctx, key)
}

func AddPublishInfoNullKey(ctx context.Context, uid int64, duration time.Duration) error {
	key := getPublishInfoNullKey(uid)
	return addNullKey(ctx, key, duration)
}

func PushVideoQueue(ctx context.Context, vid int64, maxCap int64) error {
	pipe := global.RedisClient.Pipeline()
	length, err := pipe.LLen(ctx, publishQueueKey).Result()
	if err != nil {
		return err
	}
	if length >= maxCap {
		for i := 0; i < int(length-maxCap+1); i++ {
			if _, err := pipe.RPop(ctx, publishQueueKey).Result(); err != nil {
				return err
			}
		}
	}
	if _, err = pipe.LPush(ctx, publishQueueKey, vid).Result(); err != nil {
		return err
	}
	if _, err := pipe.Exec(ctx); err != nil {
		// 报错后进行一次额外尝试
		if _, err = pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
