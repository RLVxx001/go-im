package group

import (
	"design/domain/group"
	"github.com/gorilla/websocket"
	"net/http"
)

var clients = make(map[uint][]*websocket.Conn) //消息专用
var broadcast = make(chan GroupMessage)

var clientNews = make(map[uint][]*websocket.Conn) //创建
var broadcastNew = make(chan GroupUser)           //

var clientRes = make(map[uint][]*websocket.Conn) //撤回
var broadcastRe = make(chan GroupMessage)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域请求（仅作为示例，生产环境请考虑安全性）
		return true
	},
} // 使用默认的WebSocket升级选项

type GroupRequest struct {
	Id            uint           `json:"id"`          //群id
	GroupId       string         `json:"groupId"`     //群号
	GroupName     string         `json:"groupName"`   //群名称
	GroupInform   string         `json:"groupInform"` //群公告
	GroupUsers    []GroupUser    `json:"groupUsers"`  //群用户
	GroupMessages []GroupMessage `json:"groupMessages"`
}

// GroupMessage 群消息表
type GroupMessage struct {
	Id            uint   `json:"id"`            //群消息id
	MessageOwner  uint   `json:"messageOwner"`  //消息所属用户
	MessageSender uint   `json:"messageSender"` //消息发送用户
	GroupId       uint   `json:"groupId"`       //消息接收群
	Message       string `json:"message"`       //消息
	MessageKey    uint   `json:"messageKey"`    //消息key
	IsRead        bool   `json:"isRead"`        //是否已读
}

// GroupUser 群用户表
type GroupUser struct {
	ID      uint   `json:"id"`      //群用户id
	GroupId uint   `json:"groupId"` //群表id
	UserId  uint   `json:"userId"`  //用户id
	IsAdmin uint   `json:"isAdmin"` //是否管理
	IsGag   bool   `json:"isGag"`   //是否被禁言
	Text    string `json:"text"`    //群备注
}

func ToGroup(req GroupRequest) group.Group {
	return group.Group{
		GroupId:     req.GroupId,
		GroupName:   req.GroupName,
		GroupInform: req.GroupInform,
	}
}

func ToGroupUser(req GroupUser) group.GroupUser {
	return group.GroupUser{
		ID:      req.ID,
		GroupId: req.GroupId,
		UserId:  req.UserId,
		IsAdmin: req.IsAdmin,
		IsGag:   req.IsGag,
		Text:    req.Text,
	}
}

func ToResponseGroupUser(req group.GroupUser) GroupUser {
	return GroupUser{
		ID:      req.ID,
		GroupId: req.GroupId,
		UserId:  req.UserId,
		IsAdmin: req.IsAdmin,
		IsGag:   req.IsGag,
		Text:    req.Text,
	}
}

func ToResponseGroupMessage(req group.GroupMessage) GroupMessage {
	return GroupMessage{
		Id:            req.ID,
		MessageOwner:  req.MessageOwner,
		MessageSender: req.MessageSender,
		GroupId:       req.GroupId,
		Message:       req.Message,
		MessageKey:    req.MessageKey,
		IsRead:        req.IsRead,
	}
}
