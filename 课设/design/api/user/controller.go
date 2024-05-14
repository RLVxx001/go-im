package user

import (
	"design/config"
	"design/domain/user"
	"design/utils/api_helper"
	jwtHelper "design/utils/jwt"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
		http.StatusOK, LoginResponse{
			Username: currentUser.Username,
			UserId:   currentUser.ID,
			Token:    currentUser.Token,
			Account:  currentUser.Account,
			Email:    currentUser.Email,
			Img:      currentUser.Img,
		})
}

// 验证token
func (c *Controller) VerifyToken(g *gin.Context) {
	currentUser, err := c.userService.GetById(api_helper.GetUserId(g))
	if err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	g.JSON(
		http.StatusOK, LoginResponse{
			Username: currentUser.Username,
			UserId:   currentUser.ID,
			Token:    currentUser.Token,
			Account:  currentUser.Account,
			Email:    currentUser.Email,
			Img:      currentUser.Img,
		})

}

// 上传头像
func (c *Controller) Upload(g *gin.Context) {

	if _, err := c.userService.GetById(api_helper.GetUserId(g)); err != nil {
		api_helper.HandleErrorToken(g, err)
		return
	}
	file, header, err := g.Request.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, errors.New("Failed to retrieve file"))
		return
	}
	ext := filepath.Ext(header.Filename) //获取文件后缀
	if ext != ".jpg" && ext != ".webp" && ext != ".png" {
		api_helper.HandleError(g, errors.New("文件格式不符合"))
		return
	}
	now := time.Now()
	timestampSeconds := now.Unix()

	header.Filename = fmt.Sprintf("%dto%d%s", api_helper.GetUserId(g), timestampSeconds, ext)

	// 指定上传目录，比如 "uploads"
	uploadDir := "./public/images"

	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		api_helper.HandleError(g, errors.New("Failed to create upload directory"))
		return
	}

	// 构建文件的完整路径
	filePath := filepath.Join(uploadDir, header.Filename)

	// 创建文件用于写入
	outFile, err := os.Create(filePath)
	if err != nil {
		api_helper.HandleError(g, errors.New("Failed to create file"))
		return
	}
	defer outFile.Close()

	// 复制文件内容
	if _, err := io.Copy(outFile, file); err != nil {
		api_helper.HandleError(g, errors.New("Failed to copy file"))
		return
	}
	if err := c.userService.UpdateImg(filePath, api_helper.GetUserId(g)); err != nil {
		api_helper.HandleError(g, errors.New("头像更换失败"))
		return
	}
	// 上传成功
	g.JSON(
		http.StatusOK, LoginResponse{
			Img: filePath,
		})
}
