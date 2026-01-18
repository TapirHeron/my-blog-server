package core

import (
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	addr := global.Config.System.Addr()
	router := initialize.InitRouter()
	server := initServer(addr, router)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(server.ListenAndServe().Error())
}
