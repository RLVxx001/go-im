package usertoUser

import (
	"design/domain/user"
	"design/domain/usertoUser"
	"design/utils/api_helper"
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
	if _, err := c.userServer.GetById(req.UserOwner); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	if _, err := c.userServer.GetById(req.UserTarget); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	user2 := usertoUser.NewUsertoUser(req.UserOwner, req.UserTarget, req.Remarks)
	if err := c.server.Create(user2); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusOK, ToUserResponse(user2))
}
