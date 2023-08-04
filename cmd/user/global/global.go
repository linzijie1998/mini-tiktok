package global

import (
	"github.com/linzijie1998/mini-tiktok/cmd/user/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	MongoClient *mongo.Client
)

var (
	ExpireDurationNullKey      time.Duration // 空值缓存过期时间
	ExpireDurationUserBaseInfo time.Duration // 用户基本信息缓存过期时间
)
