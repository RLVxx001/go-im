package userImg

import (
	"design/domain/userImg"
	"design/utils/api_helper"
	"design/utils/img"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	s *userImg.Service
}

// 实例化
func NewUserController(service *userImg.Service) *Controller {
	return &Controller{
		s: service,
	}
}

// 新建图片
func (c *Controller) Create(g *gin.Context) {
	filename, err := img.Create(g)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	newUserImg := userImg.NewUserImg(filename, userId)
	err = c.s.Create(newUserImg)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ToImg(*newUserImg))
}

// 查找图片
func (c *Controller) GetByUser(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	imgs, err := c.s.GetByUser(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	var responses []Img
	for _, i := range imgs {
		responses = append(responses, ToImg(i))
	}
	g.JSON(http.StatusOK, responses)
}

// 删除图片
func (c *Controller) Delete(g *gin.Context) {
	var req Img
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.s.Delete(req.Id, userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}
