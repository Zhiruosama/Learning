package main

import (
	"log"
	"minimal_chat/internal/httpserver"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 关闭默认debug日志
	gin.SetMode(gin.ReleaseMode)

	r := httpserver.BuildRouter()

	log.Println("Server listening at http://8.154.30.80:8322")
	http.ListenAndServe(":8322", r)
}
