package main

import (
	"design/api"
	"design/api/usertoUser"
	"design/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

type W struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	// 返回 JSON 数据给前端
	c.JSON(200, W{"aa", "bb"})
}

func main() {
	g := gin.Default()
	g.Use(func(c *gin.Context) {
		// 允许所有来源进行访问，这里仅作为示例，实际生产环境中应当严格限制允许的来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的方法，如 GET, POST, PUT, DELETE 等
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 预检间隔时间
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 允许携带凭证（Cookies, HTTP认证及客户端SSL证明等）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是 OPTIONS 请求，则直接返回，因为 OPTIONS 请求是用来检查服务器是否支持某个请求方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	})
	g.Use(jwtMiddleware())
	api.RegisterHandlers(g)
	{ //启动websocket辅助函数
		go usertoUser.SocketSend()
		go usertoUser.SocketCreate()
		go usertoUser.SocketRevocation()
	}
	g.Run(":8080")
}

// 中间件：检查JWT，但仅当请求不是登录或注册时
func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求是否是WebSocket升级请求
		if c.Request.Header.Get("Upgrade") == "websocket" &&
			c.Request.Header.Get("Connection") == "Upgrade" &&
			websocket.IsWebSocketUpgrade(c.Request) {
			// 如果是WebSocket升级请求，则不执行后续中间件，直接处理WebSocket连接
			fmt.Println("websocket链接.......")
			c.Set("userId", 7)
			c.Next()
			// 注意：这里的c.Next()实际上不会被执行，因为我们会在下面处理WebSocket连接
			// 你需要自定义逻辑来处理WebSocket连接
			return
		}

		// 检查请求路径是否是登录或注册
		if strings.HasPrefix(c.Request.URL.Path, "/user/login") || strings.HasPrefix(c.Request.URL.Path, "/register") {
			// 如果是登录或注册请求，则直接跳过JWT验证
			fmt.Println("跳过JWT验证")
			c.Next()
			return
		}

		// 从请求头或请求体（取决于您的JWT存储方式）中获取JWT令牌
		tokenString := c.Request.Header.Get("Authorization") // 假设JWT在Authorization头中，格式为"Bearer <token>"

		UserId, err := jwt.Decoded(tokenString)
		if err != nil {
			// 如果解析失败，返回错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		// 如果JWT验证成功，您可以将解析后的用户信息（如ID）设置到上下文中供后续使用
		c.Set("userId", UserId) // 假设您从JWT中解析出了userID
		fmt.Println("通过JWT验证")
		// 继续处理请求
		c.Next()
	}
}
