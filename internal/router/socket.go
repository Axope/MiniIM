package router

import (
	"MiniIM/internal/manager"
	"MiniIM/internal/server"
	"MiniIM/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// parme: uuid
func Socket(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		log.Logger.Error("uuid empty")
		return
	}
	if !manager.UserManager.CheckUserExistByUuid(uuid) {
		log.Logger.Error("socket parme error", zap.Any("uuid = ", uuid))
		return
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Logger.Error("socket upgrade error", zap.Any("err = ", err))
		return
	}

	client := server.NewClient(ws, uuid)
	server.RootServer.Login <- client

	go client.Read()
	go client.Write()
	go client.GroupService()
}
