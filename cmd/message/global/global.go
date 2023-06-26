package global

import (
	"github.com/linzijie1998/mini-tiktok/cmd/message/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Configs     config.ServiceConfigs
	MongoClient *mongo.Client
	Viper       *viper.Viper
)
