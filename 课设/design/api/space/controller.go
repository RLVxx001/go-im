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

func (r *Controller) FindComment(g *gin.Context) {
	var req FindCommentRequest
	g.ShouldBind(&req)
	comments, err := r.spaceService.FindComments(req.TrendId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, ToComments(comments))
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

	err = r.spaceService.CreateTrends(trend)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(
		http.StatusCreated, CreateTrendResponse{
			SpaceId: req.SpaceId,
		})
}

func (r *Controller) FindTrends(g *gin.Context) {
	var req FindTrendsRequest
	g.ShouldBind(&req)
	userId := req.UserId
	Trends, err := r.spaceService.FindTrends(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, ToFindTrendsResps(Trends))

}

func (r *Controller) FindTrend(g *gin.Context) {
	var req FindTrendRequest
	g.ShouldBind(&req)
	trend, err := r.spaceService.FindTrend(req.TrendId)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	g.JSON(200, ToFindTrendsResp(trend))

}

func (r *Controller) CreateComment(g *gin.Context) {
	var req CreateCommentRequest
	err := g.ShouldBind(&req)
	fmt.Printf("%v\n", req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	err = r.spaceService.CreateComment(req.UserId, req.Detail, req.TrendId)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	g.JSON(200, req)
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
