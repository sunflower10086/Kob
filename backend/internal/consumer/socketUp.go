package consumer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WsHandler(ctx *gin.Context, hub *Hub, userId string) {
	ServeWs(ctx, hub, userId, ctx.Writer, ctx.Request)
	log.Println("userId", userId)
}

// ServeWs handles websocket requests from the peer.
// 调用读和写的两个方法
func ServeWs(ctx *gin.Context, hub *Hub, userId string, w http.ResponseWriter, r *http.Request) {
	// 升级为webSocket协议
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), UserId: userId}
	// 注册这个连接
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump(ctx)
	go client.ReadPump(ctx)

	// TODO: 启动一个goroutine，检测是否有消息，向前端发送消息
	//go client.SendMsg()
}
