package group

import (
	wsServer "design/api/ws"
	"design/domain/group"
	"design/domain/user"
	"design/utils/api_helper"
	"design/utils/img"
	"design/utils/pagination"
	"design/utils/webSocketDecoded"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Controller struct {
	s           *group.Service
	userService *user.Service
}

func NewController(s *group.Service, userService *user.Service) *Controller {
	return &Controller{
		s:           s,
		userService: userService,
	}
}

// 个人查找所有群聊
func (c *Controller) FidGroups(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	groups, err := c.s.FidGroups(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	var gs []GroupRequest
	for i, j := range groups {
		mp := make(map[uint]user.User)
		for k, o := range j.GroupUsers {
			groupUser, err := c.userService.GetById(o.UserId)
			if err == nil {
				groups[i].GroupUsers[k].User = groupUser
				mp[o.UserId] = groupUser
			}
		}
		for k, o := range j.GroupMessages {

			u, ok := mp[o.MessageSender]
			if ok {
				groups[i].GroupMessages[k].SenderUser.User = u
			} else {
				groupUser, err := c.userService.GetById(o.MessageSender)
				if err == nil {
					groups[i].GroupMessages[k].SenderUser.User = groupUser
				}
			}
		}
		gs = append(gs, ToGroupRequest(groups[i]))
	}
	g.JSON(http.StatusOK, gs)
}

// 个人查找单群聊
func (c *Controller) FidGroup(g *gin.Context) {
	var req GroupRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	group, err := c.s.FidGroup(req.Id, userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	mp := make(map[uint]user.User)
	for k, o := range group.GroupUsers {
		user, err := c.userService.GetById(o.UserId)
		if err == nil {
			mp[o.UserId] = user
			group.GroupUsers[k].User = user
		}
	}
	for k, o := range group.GroupMessages {
		u, ok := mp[o.SenderUser.UserId]
		if ok {
			group.GroupMessages[k].SenderUser.User = u
		} else {
			group.GroupMessages[k].SenderUser.User, _ = c.userService.GetById(o.MessageSender)
		}

	}
	g.JSON(http.StatusOK, ToGroupRequest(*group))
}

// 创建群聊
func (c *Controller) CreateGroup(ws *websocket.Conn, mp map[string]interface{}, userid uint) {

	fmt.Println("验证成功：userid:", userid)

	var req GroupRequest
	err := webSocketDecoded.DecodedMap(mp, &req)
	fmt.Printf("req: %v\n", req)
	if err != nil {
		err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
		if err != nil {
			return
		}
		return
	}

	if _, err := c.userService.GetById(userid); err != nil {
		_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
		return
	}
	p := make(map[uint]bool)
	//筛选合理用户
	var groupUsers []group.GroupUser
	groupUsers = append(groupUsers, group.GroupUser{UserId: userid, IsAdmin: 2}) //创建人为群主
	p[userid] = true
	for _, j := range req.GroupUsers {
		if p[j.UserId] {
			continue
		}
		if _, err := c.userService.GetById(j.UserId); err == nil {
			groupUsers = append(groupUsers, group.GroupUser{UserId: j.UserId, IsAdmin: 1}) //其余创始人为管理
			p[j.UserId] = true
		}
	}
	group := ToGroup(req)
	group.GroupUsers = groupUsers
	err = c.s.CreateGroup(&group)
	if err != nil {
		err1 := api_helper.WsError(ws, err, "err")
		if err1 != nil {
			return
		}
		return
	}
	for _, i := range group.GroupUsers {
		wsServer.Broadcast <- wsServer.NewW(i.UserId, group, mp["event"].(string))
	}
}

// 更改群信息
func (c *Controller) UpdateGroup(g *gin.Context) {
	var req GroupRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	group := ToGroup(req)
	group.ID = req.Id
	if err := c.s.UpdateGroup(&group, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	req.GroupId = group.GroupId
	g.JSON(http.StatusOK, req)
}

// 删除群
func (c *Controller) DeleteGroup(g *gin.Context) {
	var req GroupRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}

	if err := c.s.DeleteGroup(req.Id, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 新增群用户
func (c *Controller) CreateGroupUser() {
	for {
		req := <-wsServer.GroupChan
		if _, err := c.userService.GetById(req.Owner); err != nil {
			log.Print(err)
			continue
		}
		if err := c.s.CreateGroupUser(req.Target, req.Owner, req.InviteUser); err != nil {
			log.Print(err)
			continue
		}
		wsServer.Broadcast <- wsServer.NewW(req.Owner, req, req.Event)
	}
}

// 更改群用户信息包括权限
func (c *Controller) UpdateGroupUser(g *gin.Context) {
	var req GroupUser
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	groupUser := ToGroupUser(req)
	if err := c.s.UpdateGroupUser(&groupUser, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ToResponseGroupUser(groupUser))
}

// 踢出群用户
func (c *Controller) DeleteGroupUser(g *gin.Context) {
	var req GroupUser
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	groupUser := ToGroupUser(req)
	if err := c.s.DeleteGroupUser(&groupUser, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 发送群消息
func (c *Controller) SendMessage(ws *websocket.Conn, mp map[string]interface{}, userid uint) {

	fmt.Println("验证成功：userid:", userid)

	var req GroupMessage
	err := webSocketDecoded.DecodedMap(mp, &req)
	fmt.Printf("req: %v\n", req)
	if err != nil {
		err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
		if err != nil {
			return
		}
		return
	}

	if _, err := c.userService.GetById(userid); err != nil {
		_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
		return
	}

	messages, err := c.s.SendMessage(req.GroupId, userid, req.Message, req.Img)
	if err != nil {
		err1 := api_helper.WsError(ws, err, "err")
		if err1 != nil {
			return
		}
		return
	}
	for _, i := range messages {
		response := ToResponseGroupMessage(i)
		if se, err := c.s.GetGroupUser(i.GroupId, i.MessageSender); err == nil {
			response.SenderUser = ToResponseGroupUser(*se)
		}

		user, err := c.userService.GetById(response.MessageSender)
		if err == nil {
			response.SenderUser.User.Username = user.Username
			response.SenderUser.User.Account = user.Account
			response.SenderUser.User.Img = user.Img
		}
		wsServer.Broadcast <- wsServer.NewW(i.MessageOwner, response, mp["event"].(string))
	}

}

// 撤回群消息
func (c *Controller) RevocationMessage(ws *websocket.Conn, mp map[string]interface{}, userid uint) {

	fmt.Println("验证成功：userid:", userid)

	var req GroupMessage
	err := webSocketDecoded.DecodedMap(mp, &req)
	fmt.Printf("req: %v\n", req)
	if err != nil {
		err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
		if err != nil {
			return
		}
		return
	}
	if _, err := c.userService.GetById(userid); err != nil {
		_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
		return
	}

	messages, err := c.s.RevocationMessage(req.Id, userid)
	if err != nil {
		err1 := api_helper.WsError(ws, err, "err")
		if err1 != nil {
			return
		}
		return
	}
	fmt.Printf("%v\n", messages)
	for _, i := range messages {
		wsServer.Broadcast <- wsServer.NewW(i.MessageOwner, ToResponseGroupMessage(i), mp["event"].(string))
	}

}

// 删除个人群消息
func (c *Controller) DeleteMessage(g *gin.Context) {
	var req GroupMessage
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	if err := c.s.DeleteMessage(req.Id, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 删除个人群所有消息
func (c *Controller) DeletesMessage(g *gin.Context) {
	var req GroupRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	if err := c.s.DeletesMessage(req.Id, userId); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 已读群消息
func (c *Controller) ReadMessage(g *gin.Context) {
	var req GroupRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	c.s.ReadMessage(req.Id, userId)
	g.JSON(http.StatusOK, nil)
}

// 更新群头像
func (c *Controller) UpdateImg(g *gin.Context) {
	filepath, err := img.Create(g)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	id := pagination.ParseInt(g.Request.PostFormValue("id"), 0)
	if id == 0 {
		api_helper.HandleError(g, errors.New("头像更换失败"))
		return
	}
	if err := c.s.UpdateImg(filepath, uint(id), api_helper.GetUserId(g)); err != nil {
		api_helper.HandleError(g, errors.New("头像更换失败"))
		return
	}
	// 上传成功
	g.JSON(
		http.StatusOK, GroupRequest{
			Img: filepath,
		})
}
