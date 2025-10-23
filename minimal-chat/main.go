package main

import (
	"log"
	"minimal_chat/internal/httpserver"
	"minimal_chat/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 关闭默认debug日志
	gin.SetMode(gin.ReleaseMode)
	// 初始化数据库
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN env is empty; please export a valid DSN")
	}
	if err := models.InitDB(dsn); err != nil {
		log.Fatal("init db failed: ", err)
	}
	if err := models.AutoMigrate(); err != nil {
		log.Fatal("auto migrate failed: ", err)
	}

	// 初始化路由
	r := httpserver.BuildRouter()

	log.Println("Server listening at http://8.154.30.80:8322")
	http.ListenAndServe(":8322", r)
}
