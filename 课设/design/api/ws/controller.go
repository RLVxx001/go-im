package ws

import (
	"design/utils/api_helper"
	"design/utils/webSocketDecoded"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 自定义路由负责ws分发
var Routes = make(map[string]func(*websocket.Conn, map[string]interface{}, uint))
var Clients = make(map[*websocket.Conn]uint) //ws
var Broadcast = make(chan W)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域请求（仅作为示例，生产环境请考虑安全性）
		return true
	},
} // 使用默认的WebSocket升级选项

// ws分发处理
func Ws(g *gin.Context) {
	ws, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer deleteWs(ws)
	for {
		var userid uint
		var event string
		err, s, mp := webSocketDecoded.Decoded(ws, &userid, &event)
		if err != nil {
			err = api_helper.WsError(ws, err, s)
			if err != nil {
				deleteWs(ws)
				return
			}
			continue
		}
		Clients[ws] = userid
		f, ok := Routes[event]
		if !ok {
			fmt.Println("跳过了")
			continue
		}
		f(ws, mp, userid)
	}

}

type W struct {
	userId uint
	Data   interface{} `json:"data"`
	Event  string      `json:"event"`
}

var ApplicationChan = make(chan ApplicationAccept)
var UserChan = make(chan ApplicationAccept)
var GroupChan = make(chan ApplicationAccept)

type ApplicationAccept struct {
	Owner  uint   //所属用户或群id
	Class  uint   //0表示用户申请用户 1表示用户申请群 2表示群邀请用户
	Target uint   //对方用户或群id
	Event  string //辨别
}

func NewW(userId uint, data interface{}, event string) W {
	return W{userId, data, event}
}

// 申请同意分发
func SocketApplication() {
	for {
		req := <-ApplicationChan
		fmt.Printf("%v\n", req)
		if req.Class == 0 {
			req.Event = "/usertoUser"
			UserChan <- req
		} else if req.Class == 1 {
			//
			//GroupChan <- req
		} else if req.Class == 2 {
			//i := req.Owner
			//req.Owner = req.Target
			//req.Target = i
			//GroupChan <- req
		}

	}
}

// 处理消息发送
func SocketSend() {
	//启动SocketSend
	for {
		msg := <-Broadcast
		fmt.Printf("%v\n", msg)
		for client, j := range Clients {
			if j == msg.userId {
				err := client.WriteJSON(msg)

				if err != nil {
					log.Printf("error: %v\n", err)
					deleteWs(client)
				}
			}
		}
	}

}

// 清理ws
func deleteWs(ws *websocket.Conn) {
	if _, ok := Clients[ws]; !ok {
		log.Printf("已经关闭过ws------")
		return
	}
	delete(Clients, ws)
	log.Printf("关闭ws------")
	ws.Close()
}
