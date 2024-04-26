package main

import (
	"github.com/gin-gonic/gin"
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
	g.GET("/login", Login)
	g.Run()
}
