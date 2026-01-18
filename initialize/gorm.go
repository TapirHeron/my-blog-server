package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"server/global"
)

func InitGorm() *gorm.DB {
	mysqlConfig := global.Config.Mysql
	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConfig.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to Mysql:", zap.Error(err))
		os.Exit(1)
	}
	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("Failed to get DB instance:", zap.Error(err))
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	return db
}
