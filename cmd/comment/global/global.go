package global

import (
	"github.com/linzijie1998/mini-tiktok/cmd/comment/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Configs     config.ServiceConfigs
	GormDB      *gorm.DB
	RedisClient *redis.Client
	Viper       *viper.Viper
	MongoDB     *mongo.Database
)
