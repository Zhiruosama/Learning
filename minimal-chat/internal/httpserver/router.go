package httpserver

import (
	"minimal_chat/internal/session"
	"minimal_chat/internal/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Cookie 会话
	r.Use(session.Middleware())

	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 登录/会话
	r.POST("/login", Login)
	r.GET("/me", Me)
	r.GET("/logout", Logout)

	//历史查询
	r.GET("/rooms/:id/messages", ListRoomMessages)
	r.GET("/messages/private", ListPrivateMessages)

	//WebSocket入口
	r.GET("/ws", ws.Start)

	return r
}
