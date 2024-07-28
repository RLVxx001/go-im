package api_helper

import (
	"design/utils/pagination"

	"github.com/gin-gonic/gin"
)

var userIdText = "userId"

// 从context获得用户id
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}

// 从context设置用户id
func SetUserId(g *gin.Context, id uint) {
	g.Set(userIdText, id)
}
