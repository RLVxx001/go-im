package api

import (
	groupUserApi "design/api/group"
	spaceApi "design/api/space"
	userApi "design/api/user"
	userApplicationApi "design/api/userApplication"
	userImgApi "design/api/userImg"
	usertoUserApi "design/api/usertoUser"
	"design/api/ws"
	"design/config"
	"design/domain/group"
	"design/domain/space"
	"design/domain/user"
	"design/domain/userApplication"
	"design/domain/userImg"
	"design/domain/usertoUser"
	"design/utils/database_handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

// Databases 结构体
type Databases struct {
	userRepository              *user.Repository
	usertouserRepository        *usertoUser.Repository
	usertouserMessageRepository *usertoUser.MessageRepository
	userApplicationRepository   *userApplication.Repository
	groupRepository             *group.Repository
	groupUserRepository         *group.UserRepository
	groupMessageRepository      *group.MessageRepository
	spaceRepository             *space.SpaceRepository
	trendsRepository            *space.TrendsRepository
	commentRepository           *space.CommentRepository
	userImgRepository           *userImg.Repository
	messageRepository           *space.MessageRepository
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}

// 根据配置文件创建数据库
func CreateDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	config.SecretKey = AppConfig.SecretKey
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}
	m := AppConfig.DatabaseSettings
	var dns = fmt.Sprintf("%s:%s@%s/%s?%s", m.Username, m.Password, m.DatabaseURIL, m.DatabaseName, m.DatabaseURIR)

	db := database_handler.NewMySQLDB(dns)
	log.Printf("%v", db)
	return &Databases{
		userRepository:              user.NewUserRepository(db),
		usertouserRepository:        usertoUser.NewRepository(db),
		usertouserMessageRepository: usertoUser.NewMessageRepository(db),
		userApplicationRepository:   userApplication.NewRepository(db),
		groupRepository:             group.NewRepository(db),
		groupUserRepository:         group.NewUserRepository(db),
		groupMessageRepository:      group.NewMessageRepository(db),
		spaceRepository:             space.NewSpaceRepository(db),
		trendsRepository:            space.NewTrendsRepository(db),
		commentRepository:           space.NewCommentRepository(db),
		messageRepository:           space.NewMessageRepository(db),
		userImgRepository:           userImg.NewRepository(db),
	}
}

// 注册所有控制器
func RegisterHandlers(r *gin.Engine) {

	dbs := *CreateDBs()
	RegisterUserHandlers(r, dbs)
	RegisterUsertoUserHandlers(r, dbs)
	RegisterUserApplicationHandlers(r, dbs)
	RegisterGroupHandlers(r, dbs)
	RegisterSpaceHandlers(r, dbs)
	RegisterUserImgHandlers(r, dbs)
	r.GET("/ws", ws.Ws) //注册ws
}

// 注册空间控制器
func RegisterSpaceHandlers(r *gin.Engine, dbs Databases) {
	spaceService := space.NewService(*dbs.spaceRepository, *dbs.trendsRepository, *dbs.commentRepository, *dbs.messageRepository)
	spaceController := spaceApi.NewSpaceController(spaceService, AppConfig)
	spaceGroup := r.Group("/space")
	spaceGroup.POST("/spaceadd", spaceController.CreateSpace)
	spaceGroup.POST("/fidTrends", spaceController.FindTrends)
	spaceGroup.POST("/fidTrend", spaceController.FindTrend)
	spaceGroup.POST("/fidComment", spaceController.FindComment)

	spaceGroup.POST("/addTrends", spaceController.CreateTrend)
	spaceGroup.POST("/addComment", spaceController.CreateComment)
	spaceGroup.POST("/addMessage", spaceController.CreateMessage)
	spaceGroup.POST("/fidMessage", spaceController.FindMessage)
	//spaceGroup.POST("delMessage",spaceController.)
}

// 注册用户控制器
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	userGroup.POST("/register", userController.CreateUser)
	userGroup.POST("/login", userController.Login)
	userGroup.GET("/verifyToken", userController.VerifyToken)
	userGroup.POST("/upload", userController.Upload)
	userGroup.POST("/update", userController.Update)
	userGroup.POST("/fidUser", userController.FidUser)
}

// 注册用户相册控制器
func RegisterUserImgHandlers(r *gin.Engine, dbs Databases) {
	service := userImg.NewService(*dbs.userImgRepository)
	controller := userImgApi.NewUserController(service)
	userImgGroup := r.Group("/userImg")
	userImgGroup.POST("/upload", controller.Create)
	userImgGroup.POST("/delete", controller.Delete)
	userImgGroup.GET("/getByUser", controller.GetByUser)
}

// 注册用户-用户控制器
func RegisterUsertoUserHandlers(r *gin.Engine, dbs Databases) {
	service := usertoUser.NewService(*dbs.usertouserRepository, *dbs.usertouserMessageRepository)
	userService := user.NewService(*dbs.userRepository)
	controller := usertoUserApi.NewController(service, userService)
	Group := r.Group("/usertoUser")
	{
		go controller.Create()
	}
	//RevocationWs("/usertoUser", controller.Create) 已被申请模块通过管道 特殊调用 不暴露给前端
	RevocationWs("/usertoUser/revocation", controller.Revocation)
	RevocationWs("/usertoUser/send", controller.Send)
	Group.GET("/fids", controller.Fids)
	Group.POST("/fid", controller.Fid)
	Group.POST("/update", controller.Update)
	Group.POST("/read", controller.Read)
	Group.POST("/deleteUser", controller.DeleteUser)
	Group.POST("/deleteMessage", controller.DeleteMessage)
	Group.POST("/deleteMessages", controller.DeleteMessages)
}

// 注册用户申请表
func RegisterUserApplicationHandlers(r *gin.Engine, dbs Databases) {
	service := userApplication.NewService(*dbs.userApplicationRepository)
	userService := user.NewService(*dbs.userRepository)
	groupService := group.NewService(*dbs.groupRepository, *dbs.groupMessageRepository, *dbs.groupUserRepository)
	usertoUserService := usertoUser.NewService(*dbs.usertouserRepository, *dbs.usertouserMessageRepository)
	controller := userApplicationApi.NewController(userService, groupService, usertoUserService, service)
	Group := r.Group("/userApplication")
	Group.POST("", controller.Application)
	Group.GET("/fids", controller.Fids)

}

// 注册群控制器
func RegisterGroupHandlers(r *gin.Engine, dbs Databases) {
	service := group.NewService(*dbs.groupRepository, *dbs.groupMessageRepository, *dbs.groupUserRepository)
	userService := user.NewService(*dbs.userRepository)
	controller := groupUserApi.NewController(service, userService)
	Group := r.Group("/group")
	Group.GET("/fidGroups", controller.FidGroups)
	Group.POST("/fidGroup", controller.FidGroup)
	RevocationWs("/group/createGroup", controller.CreateGroup)
	Group.POST("/updateGroup", controller.UpdateGroup)
	Group.POST("/deleteGroup", controller.DeleteGroup)
	{
		go controller.CreateGroupUser()
	}
	Group.POST("/updateGroupUser", controller.UpdateGroupUser)
	Group.POST("/deleteGroupUser", controller.DeleteGroupUser)
	RevocationWs("/group/sendMessage", controller.SendMessage)
	RevocationWs("/group/revocationMessage", controller.RevocationMessage)
	Group.POST("/deleteMessage", controller.DeleteMessage)
	Group.POST("/deletesMessage", controller.DeletesMessage)
	Group.POST("/read", controller.ReadMessage)
}

// 注册自定义路由
func RevocationWs(st string, fu func(*websocket.Conn, map[string]interface{}, uint)) {
	ws.Routes[st] = fu
}
