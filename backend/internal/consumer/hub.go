package consumer

import (
	"backend/conf/logger"
	"log"
)

// Hub maintains the set of active Clients and broadcasts messages to the
// Clients.
type Hub struct {
	// Registered Clients.
	Clients map[string]*Client

	// Inbound messages from the Clients.
	broadcast chan []byte

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		Clients:    make(map[string]*Client, 10),
	}
}

// Run hub的事件中心，有注册，注销，群发消息
func (h *Hub) Run() {
	//ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		// 进行注册
		case client := <-h.Register:
			log.Println(client.UserId, "连接成功")
			h.Clients[client.UserId] = client

			// 注销连接
		case client := <-h.unregister:
			if _, ok := h.Clients[client.UserId]; ok {
				log.Println(client.UserId, "断开连接")
				delete(h.Clients, client.UserId)
				close(client.Send)
			}

		// 广播消息
		case message := <-h.broadcast:
			logger.SugarLogger.Infof("broadcast %s", string(message))
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.UserId)
				}
			}
		}
	}
}
