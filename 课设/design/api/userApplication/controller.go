package userApplication

import (
	"design/domain/user"
	"design/domain/userApplication"
	"design/utils/api_helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	userService *user.Service
	service     *userApplication.Service
}

// 实例化
func NewController(userService *user.Service, service *userApplication.Service) *Controller {
	return &Controller{
		userService: userService,
		service:     service,
	}
}

// 创建申请
func (c *Controller) Create(g *gin.Context) {
	var req CreateRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	//校验用户是否合法
	req.UserOwner = api_helper.GetUserId(g)
	if _, err := c.userService.GetById(req.UserOwner); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	if _, err := c.userService.GetById(req.UserTarget); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	if req.IsDown { //拒绝
		if err := c.service.Update(req.UserOwner, req.UserTarget); err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, nil)
		return
	}

	newUser := userApplication.NewUserApplication(req.UserOwner, req.UserTarget, req.Remarks, req.Text)

	u, err := c.service.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, ToCreateResponse(*u))
}

// 查询申请
func (c *Controller) Fids(g *gin.Context) {
	var req CreateRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	//校验用户是否合法
	req.UserOwner = api_helper.GetUserId(g)
	if _, err := c.userService.GetById(req.UserOwner); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}

	fids, err := c.service.Fids(req.UserOwner)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	var responses []CreateResponse
	for i, j := range fids {

		responses = append(responses, ToCreateResponse(j))

		if responses[i].UserOwner != req.UserOwner {
			u, err := c.userService.GetById(responses[i].UserOwner)

			if err != nil {
				continue
			}
			responses[i].UserResponse = UserResponse{
				Username: u.Username,
				Account:  u.Account,
				Img:      u.Img,
			}
		} else {
			u, err := c.userService.GetById(responses[i].UserTarget)

			if err != nil {
				continue
			}
			responses[i].UserResponse = UserResponse{
				Username: u.Username,
				Account:  u.Account,
				Img:      u.Img,
			}
		}

	}

	g.JSON(http.StatusOK, responses)
}
