package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob("templates/**/*")

	router.Static("/static", "./static")
	//Get请求传值
	router.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	//POST传值
	router.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单POST过来的数据
	router.POST("/doAddUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
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
