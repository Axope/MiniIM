package main

import (
	"MiniIM/configs"
	"MiniIM/internal/dao/pool"
	"MiniIM/internal/router"
	"MiniIM/internal/server"
	"MiniIM/pkg/log"

	"go.uber.org/zap"
)

func main() {
	c := configs.GetConfig()

	log.InitLogger(c.Log.Level, c.Log.Path)
	defer log.Logger.Sync()
	log.Logger.Info("load configs: ", zap.Any("config", c))

	pool.Init()
	// pool.GetDB().AutoMigrate(&models.User{})

	// 启动服务
	go server.RootServer.Run()

	router := router.Create()
	router.Run(":9876")

}
