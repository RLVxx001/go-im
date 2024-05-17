package group

import (
	"design/config"
	"design/domain/group"
	"design/domain/user"
	"design/utils/api_helper"
	"design/utils/jwt"
	"design/utils/pagination"
	"fmt"
	"github.com/gin-gonic/gin"
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

// 个人查找群聊
func (c *Controller) FidGroup(g *gin.Context) {
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
	for _, i := range groups {
		gs = append(gs, ToGroupRequest(i))
	}
	g.JSON(http.StatusOK, gs)
}

// 创建群聊
func (c *Controller) CreateGroup(g *gin.Context) {
	ws, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	var userid uint
	defer deleteWs(ws, userid, clientNews)
	fmt.Println("创建链接中--------")
	for {
		var token config.Token
		err = ws.ReadJSON(&token)
		fmt.Printf("%v\n", token)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}
		if token.Type == "auth" {
			id, err := jwt.Decoded(token.Token)
			if err != nil {
				err := api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
				if err != nil {
					return
				}
				continue
			}
			userid = uint(pagination.ParseInt(id, -1))
			break
		} else {
			return
		}
	}
	fmt.Println("验证成功：userid:", userid)
	clientNews[userid] = append(clientNews[userid], ws) //添加
	for {
		var req GroupRequest
		err = ws.ReadJSON(&req)
		fmt.Printf("req: %v\n", req)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}

		if _, err := c.userService.GetById(userid); err != nil {
			_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "auth")
			return
		}
		mp := make(map[uint]bool)
		//筛选合理用户
		var groupUsers []group.GroupUser
		groupUsers = append(groupUsers, group.GroupUser{UserId: userid, IsAdmin: 2}) //创建人为群主
		mp[userid] = true
		for _, j := range req.GroupUsers {
			if mp[j.UserId] {
				continue
			}
			if _, err := c.userService.GetById(j.UserId); err == nil {
				groupUsers = append(groupUsers, group.GroupUser{UserId: j.UserId})
				mp[j.UserId] = true
			}
		}
		group := ToGroup(req)
		group.GroupUsers = groupUsers
		err := c.s.CreateGroup(&group)
		if err != nil {
			err1 := api_helper.WsError(ws, err, "")
			if err1 != nil {
				return
			}
			continue
		}

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
func (c *Controller) CreateGroupUser(g *gin.Context) {
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
	//筛选合理用户
	mp := make(map[uint]bool)
	for _, j := range req.GroupUsers {
		if mp[j.UserId] {
			continue
		}
		if _, err := c.userService.GetById(j.UserId); err == nil {
			if c.s.CreateGroupUser(req.Id, j.UserId, userId) == nil {

			}
			mp[j.UserId] = true
		}
	}
	g.JSON(http.StatusOK, nil)

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
func (c *Controller) SendMessage(g *gin.Context) {
	ws, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	var userid uint
	defer deleteWs(ws, userid, clientNews)
	fmt.Println("发送链接中--------")
	for {
		var token config.Token
		err = ws.ReadJSON(&token)
		fmt.Printf("%v\n", token)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}
		if token.Type == "auth" {
			id, err := jwt.Decoded(token.Token)
			if err != nil {
				err := api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
				if err != nil {
					return
				}
				continue
			}
			userid = uint(pagination.ParseInt(id, -1))
			break
		} else {
			return
		}
	}
	fmt.Println("验证成功：userid:", userid)
	clients[userid] = append(clients[userid], ws) //添加

	for {
		var req GroupMessage
		err = ws.ReadJSON(&req)
		fmt.Printf("req: %v\n", req)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}

		if _, err := c.userService.GetById(userid); err != nil {
			_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "auth")
			return
		}

		messages, err := c.s.SendMessage(req.GroupId, userid, req.Message)
		if err != nil {
			err1 := api_helper.WsError(ws, err, "")
			if err1 != nil {
				return
			}
			continue
		}
		for _, i := range messages {
			broadcast <- ToResponseGroupMessage(i)
		}
	}

}

// 撤回群消息
func (c *Controller) RevocationMessage(g *gin.Context) {
	ws, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	var userid uint
	defer deleteWs(ws, userid, clientRes)
	fmt.Println("撤回链接中--------")
	for {
		var token config.Token
		err = ws.ReadJSON(&token)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}
		if token.Type == "auth" {
			id, err := jwt.Decoded(token.Token)
			if err != nil {
				err := api_helper.WsError(ws, api_helper.ErrInvalidToken, "token")
				if err != nil {
					return
				}
				continue
			}
			userid = uint(pagination.ParseInt(id, -1))
			break
		} else {
			return
		}
	}

	fmt.Println("验证成功：userid:", userid)
	clientRes[userid] = append(clientRes[userid], ws) //添加
	for {
		var req GroupMessage
		err = ws.ReadJSON(&req)
		fmt.Printf("req: %v\n", req)
		if err != nil {
			err := api_helper.WsError(ws, api_helper.ErrInvalidBody, "auth")
			if err != nil {
				return
			}
			continue
		}
		if _, err := c.userService.GetById(userid); err != nil {
			_ = api_helper.WsError(ws, api_helper.ErrInvalidToken, "auth")
			return
		}

		messages, err := c.s.RevocationMessage(req.Id, userid)
		if err != nil {
			err1 := api_helper.WsError(ws, err, "")
			if err1 != nil {
				return
			}
			continue
		}
		fmt.Printf("%v\n", messages)
		for _, i := range messages {
			broadcastRe <- ToResponseGroupMessage(i)
		}
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
