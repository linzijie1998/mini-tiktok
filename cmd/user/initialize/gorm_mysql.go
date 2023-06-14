package initialize

import (
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMySQL() (*gorm.DB, error) {
	m := global.Configs.MySQL
	if m.DBName == "" {
		return nil, errno.ServiceErr.WithMessage("配置文件中未指定数据库名")
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DSN(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
