package group

import (
	"github.com/gorilla/websocket"
	"log"
)

// 处理消息发送
func SocketSend() {
	//启动SocketSend
	for {
		msg := <-broadcast

		for _, client := range clients[msg.MessageOwner] {
			err := client.WriteJSON(msg)

			if err != nil {
				log.Printf("error: %v\n", err)
				deleteWs(client, msg.MessageOwner, clients)
			}
		}
	}

}

// 处理创建用户-用户
func SocketCreate() {
	for {
		msg := <-broadcastNew
		for _, client := range clientNews[msg.UserId] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v\n", err)
				deleteWs(client, msg.UserId, clientNews)
			}
		}
	}
}

// 处理撤回消息请求
func SocketRevocation() {
	for {
		msg := <-broadcastRe

		for _, client := range clientRes[msg.MessageOwner] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v\n", err)
				deleteWs(client, msg.MessageOwner, clientRes)
			}
		}
	}
}

func deleteWs(ws *websocket.Conn, id uint, clients map[uint][]*websocket.Conn) {
	for index, i := range clients[id] {
		if i == ws {
			clients[id] = append(clients[id][:index], clients[id][index+1:]...) //删除
		}
	}
	log.Printf("关闭ws------")
	ws.Close()
}
