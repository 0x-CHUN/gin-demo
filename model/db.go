package model

import (
	"gin-demo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Database(conn string) {
	LOG := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		})
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: LOG,
	})
	if conn == "" || err != nil {
		utils.Log().Error("mysql lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		utils.Log().Error("mysql lost: %v", err)
		panic(err)
	}
	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
}
