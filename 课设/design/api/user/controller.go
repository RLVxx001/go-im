package user

import (
	"design/config"
	"design/domain/user"
	"design/utils/api_helper"
	"design/utils/img"
	jwtHelper "design/utils/jwt"
	"design/utils/redisYz"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

// 实例化
func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

// CreateUser godoc
// @Summary 根据给定的用户名和密码创建用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param CreateUserRequest body CreateUserRequest true "user information"
// @Success 201 {object} CreateUserResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user [post]
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}

	newUser := user.NewUser(req.Username, req.Email, req.Password, req.Password2, req.Email)
	err := c.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateUserResponse{
			Username: req.Username,
		})
}

// Login godoc
// @Summary 根据用户名和密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "user information"
// @Success 200 {object} LoginResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/login [post]
func (c *Controller) Login(g *gin.Context) {

	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	code, _ := redisYz.GetVerificationCode()
	if code != req.Yz {
		api_helper.HandleError(g, errors.New("验证码错误"))
		return
	}
	currentUser, err := c.userService.CheckUser(req.Text, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	decodedClaims := jwtHelper.VerifyToken(currentUser.Token, c.appConfig.SecretKey)
	if decodedClaims == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp": time.Now().Add(
					24 *
						time.Hour).Unix(),
				"isAdmin": currentUser.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		currentUser.Token = token
		err = c.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}

	g.JSON(
		http.StatusOK, ToLoginResponse(currentUser))
}

// 查询个人信息
func (c *Controller) GetUser(g *gin.Context) {
	userId := api_helper.GetUserId(g)
	user, err := c.userService.GetById(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ToLoginResponse(user))
}

// 验证token
func (c *Controller) VerifyToken(g *gin.Context) {
	currentUser, err := c.userService.GetById(api_helper.GetUserId(g))
	if err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	g.JSON(http.StatusOK, ToLoginResponse(currentUser))

}

// 查找用户（通过账号）
func (c *Controller) FidUser(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	_, err := c.userService.GetById(api_helper.GetUserId(g))
	if err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	getUser, err := c.userService.GetUser(req.Text)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ToLoginResponse(getUser))
}

// 更新信息
func (c *Controller) Update(g *gin.Context) {
	var req LoginResponse
	fmt.Printf("%v\n", req)
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, api_helper.ErrInvalidBody)
		return
	}
	userid := api_helper.GetUserId(g)
	user1, err := c.userService.GetById(userid)
	if err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	User2 := user.User{
		Account:  req.Account,
		Signed:   req.Signed,
		Birthday: req.Birthday,
	}
	User2.ID = user1.ID
	if err := c.userService.UpdateUser(&User2); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ToLoginResponse(User2))
}

// 上传头像
func (c *Controller) Upload(g *gin.Context) {
	if _, err := c.userService.GetById(api_helper.GetUserId(g)); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	filepath, err := img.Create(g)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	if err := c.userService.UpdateImg(filepath, api_helper.GetUserId(g)); err != nil {
		api_helper.HandleError(g, errors.New("头像更换失败"))
		return
	}
	// 上传成功
	g.JSON(
		http.StatusOK, LoginResponse{
			Img: filepath,
		})
}

// 获取验证码
func (c *Controller) CreateYz(g *gin.Context) {
	_ = redisYz.SetVerificationCode()
	code, _ := redisYz.GetVerificationCode()
	g.JSON(http.StatusOK, CreateUserRequest{Yz: code})
}
