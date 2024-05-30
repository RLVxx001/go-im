package space

import (
	"design/config"
	"design/domain/space"
	"design/utils/api_helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	spaceService *space.Service
	appConfig    *config.Configuration
}

func NewSpaceController(service *space.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		spaceService: service,
		appConfig:    appConfig,
	}
}

func (r *Controller) CreateSpace(g *gin.Context) {
	var req CreateSpaceResp
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = r.spaceService.CreateSpace(req.UserId)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
}

func (r *Controller) CreateTrend(g *gin.Context) {
	var req CreateTrendRequest
	err := g.ShouldBind(&req)
	fmt.Printf("%v", req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	trend := ToSpaceTrend(req)

	err = r.spaceService.CreateTrends(&trend)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusCreated, CreateTrendResponse{
			SpaceId: req.SpaceId,
		})
}

func (r *Controller) FindTrend(g *gin.Context) {
	var req FindTrendRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	var Trends []space.SpaceTrends
	Trends, err = r.spaceService.FindTrends(req.UserId)
	if err != nil {
		api_helper.HandleError(g, err)
	}
	g.JSON(200, FindTrendResponse{Trends: Trends})

}

func (r *Controller) CreateComment(g *gin.Context) {
	var req CreateCommentRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = r.spaceService.CreateComment(req.Comment)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = r.spaceService.AddComment(req.Comment)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	g.JSON(200, CreateCommentRequest{Comment: req.Comment})
}

func (r *Controller) DeleteTrends(g *gin.Context) {
	var req DeleteTrendRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = r.spaceService.DeleteTrends(req.Trend)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
	}
	g.JSON(200, DeleteTrendResponse{req.Trend})
}
