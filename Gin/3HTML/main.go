package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

// 时间戳转换为日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

func main() {
	router := gin.Default()

	//自定义模板函数 要把这个函数放在加载模板之前
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	//加载模板
	router.LoadHTMLGlob("templates/**/*")

	router.Static("/static", "./static")
	//前台
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"hobby": []string{"吃饭", "睡觉", "打游戏"},
			"news": &Article{
				Title:   "标题",
				Content: "内容",
			},
		})
	})
	router.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻首页",
		})
	})
	//后台
	router.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	router.Run()
}
