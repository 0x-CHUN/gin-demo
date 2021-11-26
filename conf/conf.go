package conf

import (
	"gin-demo/cache"
	"gin-demo/model"
	"gin-demo/utils"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		utils.Log().Panic("Load config file error ", err)
	}
	utils.BuildLogger(os.Getenv("LOG_LEVEL"))
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		utils.Log().Panic("翻译文件加载失败", err)
	}
	model.Database(os.Getenv("MYSQL"))
	cache.Redis()
}
