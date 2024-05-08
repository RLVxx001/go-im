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

//// 中间件：检查JWT，但仅当请求不是登录或注册时
//func jwtMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 检查请求路径是否是登录或注册
//		if strings.HasPrefix(c.Request.URL.Path, "/login") || strings.HasPrefix(c.Request.URL.Path, "/register") {
//			// 如果是登录或注册请求，则直接跳过JWT验证
//			c.Next()
//			return
//		}
//
//		// 从请求头或请求体（取决于您的JWT存储方式）中获取JWT令牌
//		tokenString := c.Request.Header.Get("Authorization") // 假设JWT在Authorization头中，格式为"Bearer <token>"
//		if tokenString == "" {
//			// 如果没有找到JWT，返回错误
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "JWT token missing"})
//			return
//		}
//
//		// 移除"Bearer "前缀（如果存在）
//		const bearerPrefix = "Bearer "
//		if len(tokenString) > len(bearerPrefix) && strings.EqualFold(tokenString[:len(bearerPrefix)], bearerPrefix) {
//			tokenString = tokenString[len(bearerPrefix):]
//		}
//
//		// 解析JWT令牌
//		token, err := parseJWT(tokenString)
//		if err != nil {
//			// 如果解析失败，返回错误
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
//			return
//		}
//
//		// 如果JWT验证成功，您可以将解析后的用户信息（如ID）设置到上下文中供后续使用
//		// c.Set("userId", userID) // 假设您从JWT中解析出了userID
//
//		// 继续处理请求
//		c.Next()
//	}
//}
