package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	router := gin.Default()
	//意思就是加载templates文件夹下的html文件
	//配置模板文件路径
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "首页")
	})
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "hello gin",
		})
	})
	router.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "jsonp标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})
	router.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})
	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面",
		})
	})
	router.Run()
}
