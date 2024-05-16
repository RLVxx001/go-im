package group

import (
	"design/domain/group"
	"design/domain/user"
	"design/utils/api_helper"
	"github.com/gin-gonic/gin"
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

// 创建群聊
func (c *Controller) CreateGroup(g *gin.Context) {
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
	var groupUsers []group.GroupUser
	groupUsers = append(groupUsers, group.GroupUser{UserId: userId, IsAdmin: 2}) //创建人为群主
	mp := make(map[uint]bool)
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
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
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
	messages, err := c.s.SendMessage(req.GroupId, userId, req.Message)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, messages)
}

// 撤回群消息
func (c *Controller) RevocationMessage(g *gin.Context) {
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
	messages, err := c.s.RevocationMessage(req.Id, userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, messages)
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
