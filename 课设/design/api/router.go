package api

import (
	groupUserApi "design/api/group"
	spaceApi "design/api/space"
	userApi "design/api/user"
	userApplicationApi "design/api/userApplication"
	usertoUserApi "design/api/usertoUser"
	"design/config"
	"design/domain/group"
	"design/domain/space"
	"design/domain/user"
	"design/domain/userApplication"
	"design/domain/usertoUser"
	"design/utils/database_handler"
	"fmt"
	"github.com/gin-gonic/gin"
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
}

// 注册空间控制器
func RegisterSpaceHandlers(r *gin.Engine, dbs Databases) {
	spaceService := space.NewService(*dbs.spaceRepository, *dbs.trendsRepository, *dbs.commentRepository)
	spaceController := spaceApi.NewSpaceController(spaceService, AppConfig)
	spaceGroup := r.Group("/space")
	spaceGroup.POST("/fidTrends", spaceController.FindTrend)
	spaceGroup.POST("/addTrends", spaceController.CreateTrend)
	spaceGroup.POST("/addComment", spaceController.CreateComment)
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

// 注册用户-用户控制器
func RegisterUsertoUserHandlers(r *gin.Engine, dbs Databases) {
	service := usertoUser.NewService(*dbs.usertouserRepository, *dbs.usertouserMessageRepository)
	userService := user.NewService(*dbs.userRepository)
	controller := usertoUserApi.NewController(service, userService)
	Group := r.Group("/usertoUser")
	Group.GET("", controller.Create)
	Group.GET("/revocation", controller.Revocation)
	Group.GET("/send", controller.Send)
	Group.GET("/fid", controller.Fids)
	Group.POST("/update", controller.Update)
	Group.POST("/read", controller.Read)
	Group.POST("/delete", controller.Delete)
	Group.POST("/deletes", controller.Deletes)
}

// 注册用户-用户申请表
func RegisterUserApplicationHandlers(r *gin.Engine, dbs Databases) {
	service := userApplication.NewService(*dbs.userApplicationRepository)
	userService := user.NewService(*dbs.userRepository)
	controller := userApplicationApi.NewController(userService, service)
	Group := r.Group("/userApplication")
	Group.POST("", controller.Create)
	Group.GET("/fid", controller.Fids)

}

// 注册群控制器
func RegisterGroupHandlers(r *gin.Engine, dbs Databases) {
	service := group.NewService(*dbs.groupRepository, *dbs.groupMessageRepository, *dbs.groupUserRepository)
	userService := user.NewService(*dbs.userRepository)
	controller := groupUserApi.NewController(service, userService)
	Group := r.Group("/group")
	Group.GET("/fidGroup", controller.FidGroup)
	Group.GET("/createGroup", controller.CreateGroup)
	Group.POST("/updateGroup", controller.UpdateGroup)
	Group.POST("/deleteGroup", controller.DeleteGroup)
	Group.POST("/createGroupUser", controller.CreateGroupUser)
	Group.POST("/updateGroupUser", controller.UpdateGroupUser)
	Group.POST("/deleteGroupUser", controller.DeleteGroupUser)
	Group.GET("/sendMessage", controller.SendMessage)
	Group.GET("/revocationMessage", controller.RevocationMessage)
	Group.POST("/deleteMessage", controller.DeleteMessage)
	Group.POST("/deletesMessage", controller.DeletesMessage)
}
