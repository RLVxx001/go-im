package userApplication

import (
	"design/domain/group"
	"design/domain/user"
	"design/domain/userApplication"
	"design/domain/usertoUser"
	"design/utils/api_helper"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	userService       *user.Service
	groupService      *group.Service
	usertoUserService *usertoUser.Service
	service           *userApplication.Service
}

// 实例化
func NewController(userService *user.Service, groupService *group.Service, usertoUserService *usertoUser.Service, service *userApplication.Service) *Controller {
	return &Controller{
		userService:       userService,
		groupService:      groupService,
		usertoUserService: usertoUserService,
		service:           service,
	}
}

func (c *Controller) Application(g *gin.Context) {
	var req CreateRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	if req.Stats == 0 {
		c.Create(g, req)
	} else if req.Stats == 1 {
		c.Refuse(g, req)
	} else if req.Stats == 2 {
		c.Accept(g, req)
	}
}

// 创建申请
func (c *Controller) Create(g *gin.Context, req CreateRequest) {

	application := userApplication.NewUserApplication(req.UserOwner, req.Class, req.Target, req.Remarks, req.Text)
	if application.Class == 0 { //自己申请好友
		application.UserOwner = api_helper.GetUserId(g)
		_, err := c.usertoUserService.Fid(application.UserOwner, application.Target) //先查询俩人是否是好友
		if err == nil || application.UserOwner == application.Target {
			api_helper.HandleError(g, errors.New("您已添加该好友"))
			return
		}
		create, err := c.service.Create(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, ToCreateResponse(*create))
	} else if application.Class == 1 { //自己申请群
		application.UserOwner = api_helper.GetUserId(g)
		_, err := c.groupService.GetById(application.Target) //查询群聊是否存在
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		_, err = c.groupService.GetGroupUser(application.Target, application.UserOwner) //查询群内是否有该用户
		if err == nil {
			api_helper.HandleError(g, errors.New("您已加入该群聊"))
			return
		}
		create, err := c.service.Create(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, ToCreateResponse(*create))
	} else if application.Class == 2 { //身为群管理去邀请用户
		userId := api_helper.GetUserId(g)
		groupUser, err := c.groupService.GetGroupUser(application.UserOwner, userId) //查询群内是否有该用户
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		if groupUser.IsAdmin == 0 {
			api_helper.HandleError(g, errors.New("对不起您的权限不够"))
			return
		}
		_, err = c.userService.GetById(application.Target)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		_, err = c.groupService.GetGroupUser(application.UserOwner, application.Target) //查询邀请的用户是否存在群中
		if err == nil {
			api_helper.HandleError(g, errors.New("该用户已加入该群聊"))
			return
		}
		application.InviteUser = userId //设置邀请人
		create, err := c.service.Create(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, ToCreateResponse(*create))
	}
}

// 拒绝申请
func (c *Controller) Refuse(g *gin.Context, req CreateRequest) {
	application := userApplication.NewUserApplication(req.UserOwner, req.Class, req.Target, req.Remarks, req.Text)
	if application.Class == 0 { //拒绝好友申请
		_, err := c.usertoUserService.Fid(application.UserOwner, api_helper.GetUserId(g)) //先查询俩人是否是好友
		if err == nil {
			api_helper.HandleError(g, errors.New("您已添加该好友"))
			return
		}
		application.Target = api_helper.GetUserId(g)
		err = c.service.Refuse(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, nil)
	} else if application.Class == 1 { //群管理拒绝用户
		userId := api_helper.GetUserId(g)
		groupUser, err := c.groupService.GetGroupUser(application.Target, userId) //查询群内是否有本用户
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		if groupUser.IsAdmin == 0 {
			api_helper.HandleError(g, errors.New("对不起您的权限不够"))
			return
		}
		_, err = c.groupService.GetGroupUser(application.Target, application.UserOwner) //查询群内是否存在申请用户
		if err == nil {
			api_helper.HandleError(g, errors.New("该用户已加入该群聊"))
			return
		}
		application.InviteUser = api_helper.GetUserId(g) //设置批阅人
		err = c.service.Refuse(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, nil)
	} else if application.Class == 2 { //拒绝群邀请
		userId := api_helper.GetUserId(g)
		_, err := c.groupService.GetById(application.UserOwner)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		_, err = c.groupService.GetGroupUser(application.UserOwner, userId) //查询群内是否有该用户
		if err == nil {
			api_helper.HandleError(g, errors.New("您已加入该群聊"))
			return
		}
		application.Target = api_helper.GetUserId(g)
		err = c.service.Refuse(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, nil)
	}
}

// 同意申请
func (c *Controller) Accept(g *gin.Context, req CreateRequest) {
	application := userApplication.NewUserApplication(req.UserOwner, req.Class, req.Target, req.Remarks, req.Text)
	userId := api_helper.GetUserId(g)
	if application.Class == 0 { //对于好友申请同意
		application.Target = userId
		_, err := c.userService.GetById(application.UserOwner)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		_, err = c.usertoUserService.Fid(application.UserOwner, application.Target)
		if err == nil {
			api_helper.HandleError(g, errors.New("您已添加该好友"))
			return
		}
		err = c.service.Accept(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, req)
	} else if application.Class == 1 { //身为管理员 同意用户申请
		groupUser, err := c.groupService.GetGroupUser(application.Target, userId) //查询群内是否有本用户
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		if groupUser.IsAdmin == 0 {
			api_helper.HandleError(g, errors.New("对不起您的权限不够"))
			return
		}
		_, err = c.groupService.GetGroupUser(application.Target, application.UserOwner) //查询群内是否存在申请用户
		if err == nil {
			api_helper.HandleError(g, errors.New("该用户已加入该群聊"))
			return
		}
		application.InviteUser = api_helper.GetUserId(g) //设置批阅人
		err = c.service.Accept(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, req)
	} else if application.Class == 2 { //身为用户 同意群管理发来的请求
		application.Target = userId
		_, err := c.groupService.GetById(application.UserOwner) //查询群聊
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		_, err = c.groupService.GetGroupUser(application.UserOwner, application.Target) //查询群内是否有本用户
		if err == nil {
			api_helper.HandleError(g, errors.New("您已加入该群聊"))
			return
		}
		err = c.service.Accept(application)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
		g.JSON(http.StatusOK, nil)
	}
}

// 查询申请
func (c *Controller) Fids(g *gin.Context) {

	//校验用户是否合法
	userId := api_helper.GetUserId(g)
	if _, err := c.userService.GetById(userId); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}

	fids, err := c.service.Fids(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	var responses []CreateResponse
	for i, j := range fids {

		responses = append(responses, ToCreateResponse(j))

		if j.Class == 0 {
			if j.UserOwner != userId {
				useres, err := c.userService.GetById(j.UserOwner)
				if err != nil {
					continue
				}
				responses[i].UserResponse = UserResponse{
					Username: useres.Username,
					Account:  useres.Account,
					Img:      useres.Img,
				}
			} else {
				useres, err := c.userService.GetById(j.Target)
				if err != nil {
					continue
				}
				responses[i].UserResponse = UserResponse{
					Username: useres.Username,
					Account:  useres.Account,
					Img:      useres.Img,
				}
			}

		} else if j.Class == 1 {
			groupes, err := c.groupService.GetById(j.Target)
			if err != nil {
				continue
			}
			responses[i].GroupResponse = GroupResponse{
				GroupId:   groupes.GroupId,
				GroupName: groupes.GroupName,
				Img:       "",
			}
		} else if j.Class == 2 {
			groupes, err := c.groupService.GetById(j.UserOwner)
			if err != nil {
				continue
			}
			responses[i].GroupResponse = GroupResponse{
				GroupId:   groupes.GroupId,
				GroupName: groupes.GroupName,
				Img:       "",
			}
		}

	}

	g.JSON(http.StatusOK, responses)
}
