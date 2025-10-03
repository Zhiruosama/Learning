package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "默认页面")
	})
	router.GET("/news", func(c *gin.Context) {
		c.String(200, "这是新闻页面 111")
	})
	router.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个post请求")
	})
	router.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个PUT请求")
	})
	router.DELETE("/del", func(c *gin.Context) {
		c.String(200, "这是一个DEL请求")
	})
	router.Run()
}
