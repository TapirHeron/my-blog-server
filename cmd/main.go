package main

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectES()
	initialize.InitOther()
	initialize.InitCron()
	defer func(Redis *redis.Client) {
		err := Redis.Close()
		if err != nil {
			global.Log.Error("Failed to close Redis connection:", zap.Error(err))
			os.Exit(1)
		}
	}(&global.Redis)

	core.RunServer()
}
