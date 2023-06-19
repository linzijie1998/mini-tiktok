package global

import (
	"github.com/linzijie1998/mini-tiktok/cmd/user/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	Viper       *viper.Viper
)

var (
	ExpireDurationNullKey       time.Duration
	ExpireDurationUserBaseInfo  time.Duration
	ExpireDurationVideoBaseInfo time.Duration
)
