package server

import (
	"MiniIM/internal/mq/RabbitMQ"
	"MiniIM/internal/protocol"
	"MiniIM/pkg/log"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	ClientManager *sync.Map // key: uuid, value: *Client
	Login         chan *Client
	Logout        chan *Client
	Event         chan *protocol.Message
}

var RootServer = NewServer()

func NewServer() *Server {
	return &Server{
		ClientManager: &sync.Map{},
		Login:         make(chan *Client),
		Logout:        make(chan *Client),
		Event:         make(chan *protocol.Message),
	}
}

func (server *Server) Run() {
	defer func() {
		// log.Logger.Error("Server is G")
	}()
	log.Logger.Info("Server is running")

	for {
		select {
		case c := <-server.Login:
			log.Logger.Info("client login", zap.Any("uuid", c.UserUuid))
			server.ClientManager.Store(c.UserUuid, c)

		case c := <-server.Logout:
			log.Logger.Info("client logout", zap.Any("uuid", c.UserUuid))
			server.ClientManager.Delete(c.UserUuid)
			// if _, ok := server.ClientManager.Load(c.UserUuid); ok {
			// 	server.ClientManager.Delete(c.UserUuid)
			// }

		case msg := <-server.Event:
			if msg.GetType() == protocol.MessageType_ToUser {
				log.Logger.Debug("to user", zap.Any("msg", msg.GetContent()))
				recv, ok := server.ClientManager.Load(msg.GetTo())
				if ok {
					data, err := proto.Marshal(msg)
					if err != nil {
						log.Logger.Error("send message error")
						return
					}
					if c, ok := recv.(*Client); ok {
						c.Send <- data
					}
				}
			} else if msg.GetType() == protocol.MessageType_ToGroup {
				data, err := proto.Marshal(msg)
				if err != nil {
					log.Logger.Error(err.Error())
					continue
				}
				err = RabbitMQ.Publish(msg.To, data)
				if err != nil {
					log.Logger.Error(err.Error())
				}
			} else {
				log.Logger.Debug("to group")
			}

		}

	}
}
