package usertoUser

import (
	"design/domain/user"
	"design/domain/usertoUser"
	"design/utils/api_helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	server     *usertoUser.Service
	userServer *user.Service
}

// 实例化
func NewController(server *usertoUser.Service) *Controller {
	return &Controller{server: server}
}

// 创建用户-用户链接
func (c *Controller) Create(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	//校验用户id是否正确
	req.UserOwner = api_helper.GetUserId(g)
	fmt.Printf("%v\n", req)

	if _, err := c.userServer.GetById(req.UserOwner); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	if _, err := c.userServer.GetById(req.UserTarget); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	//创建链接
	user2, err := c.server.Create(usertoUser.NewUsertoUser(req.UserOwner, req.UserTarget, req.Remarks))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, ToUserResponse(user2))
}

// 发送消息
func (c *Controller) Send(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)
	if err := c.server.Send(utou, req.Massage); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, nil)
}

// 修改用户-用户信息
func (c *Controller) Update(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)

	if err := c.server.Update(utou); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 撤回消息
func (c *Controller) Revocation(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	if len(req.UserMassages) == 0 {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)
	fidutou, err := c.server.Fid(utou.UserOwner, utou.UserTarget)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	utou.ID = fidutou.ID
	if err := c.server.Revocation(utou); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// 用户单方面删除消息
func (c *Controller) Delete(g *gin.Context) {
	var req UserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	if len(req.UserMassages) == 0 {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	req.UserOwner = api_helper.GetUserId(g)
	utou := ToUsertoUser(req)
	fidutou, err := c.server.Fid(utou.UserOwner, utou.UserTarget)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	utou.ID = fidutou.ID
	if err := c.server.DeleteMessage(utou); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}
