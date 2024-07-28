package api_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 错误处理
func HandleError(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	g.Abort()
	return
}

// 错误处理
func HandleErrorToken(g *gin.Context, err error) {

	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Type:    "token",
		})
	g.Abort()
	return
}

// ws错误处理
func WsError(ws *websocket.Conn, err error, pe string) error {

	err1 := ws.WriteJSON(ErrorResponse{
		Message: err.Error(),
		Type:    pe,
	})
	if err1 != nil {
		return err1
	}
	return nil
}
