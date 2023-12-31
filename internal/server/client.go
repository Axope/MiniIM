package server

import (
	"MiniIM/internal/manager"
	"MiniIM/internal/mq/RabbitMQ"
	"MiniIM/internal/protocol"
	"MiniIM/pkg/log"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	UserUuid string
}

func NewClient(conn *websocket.Conn, userUuid string) *Client {
	return &Client{
		Conn:     conn,
		Send:     make(chan []byte),
		UserUuid: userUuid,
	}
}

func (c *Client) Read() {
	defer func() {
		log.Logger.Info("client read close", zap.Any("uuid", c.UserUuid))
		// 通知RootServer用户登出
		RootServer.Logout <- c
		c.Conn.Close()
		close(c.Send)
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		// 这里没读到消息就是ws被关闭
		if err != nil {
			log.Logger.Error("client read message error",
				zap.Any("err = ", err),
				zap.Any("client uuid = ", c.UserUuid),
			)
			break
		}

		msg := &protocol.Message{}
		err = proto.Unmarshal(p, msg)
		if err != nil {
			log.Logger.Error(err.Error())
			continue
		}
		log.Logger.Sugar().Debug("decode msg: ", msg, "type: ", msg.GetType())

		RootServer.Event <- msg
	}
}

func (c *Client) Write() {
	defer func() {
		log.Logger.Info("client write close", zap.Any("uuid", c.UserUuid))
		c.Conn.Close()
		// 这里 Conn.Close() 会导致 Read() 中的ReadMessage() 错误
		// 让 Read() 去关闭 Send
	}()

	for msg := range c.Send {
		log.Logger.Sugar().Debug("write msg: ", msg)
		{
			tmpMsg := &protocol.Message{}
			if err := proto.Unmarshal(msg, tmpMsg); err != nil {
				log.Logger.Sugar().Debugf("debug try decode err", err)
			} else {
				log.Logger.Sugar().Debugf("write msg(decode): ", tmpMsg)
			}
		}
		if err := c.Conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
			log.Logger.Error("client write message error",
				zap.Any("err", err),
				zap.Any("client uuid", c.UserUuid),
			)
			break
		}
	}
}

func (c *Client) GroupService() {
	groupUuids, err := manager.GroupMemberManager.GetGroupsByUuid(c.UserUuid)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	if len(groupUuids) == 0 {
		log.Logger.Debug("not groups")
		return
	}
	q, err := RabbitMQ.NewQueueAndBind(c.UserUuid, groupUuids)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	msgs, err := RabbitMQ.NewConsume(q.Name)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	// 收到的数据是字节流
	for m := range msgs {
		// m.ContentType == "application/octet-stream"
		c.Send <- m.Body
	}
}
