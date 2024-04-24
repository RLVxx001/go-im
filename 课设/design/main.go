package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "-------------")
}

func main() {
	g := gin.Default()
	g.GET("/login", Login)
	g.Run()
}
