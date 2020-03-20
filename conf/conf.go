package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/util"
	"os"

	"github.com/joho/godotenv"
	"github.com/liudng/godump"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("/go/src/giligili/conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}
	a := os.Getenv("GIN_MODE")
	godump.Dump(a)
	os.Exit(1)
	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
