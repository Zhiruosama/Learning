package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Who are you")
	// })

	// r.GET("/user/:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	ctx.String(http.StatusOK, "hello %s", name)
	// })

	// r.POST("/form", func(ctx *gin.Context) {
	// 	username := ctx.PostForm("username")
	// 	password := ctx.DefaultPostForm("password", "00000")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })

	// MAP参数（字典参数）
	// r.POST("/post", func(ctx *gin.Context) {
	// 	ids := ctx.QueryMap("ids")
	// 	names := ctx.PostFormMap("names")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"ids":   ids,
	// 		"names": names,
	// 	})
	// })

	//重定向
	// r.GET("/redirect", func(ctx *gin.Context) {
	// 	ctx.Redirect(http.StatusMovedPermanently, "/index")
	// })
	// r.GET("/goindex", func(ctx *gin.Context) {
	// 	ctx.Request.URL.Path = "/"
	// 	r.HandleContext(ctx)
	// })

	// 分组路由
	// defaultHandler := func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"path": c.FullPath(),
	// 	})
	// }
	// //group: v1
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/posts", defaultHandler)
	// 	v1.GET("/series", defaultHandler)
	// }
	// // group: v2
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/posts", defaultHandler)
	// 	v2.GET("/series", defaultHandler)
	// }

	//文件上传
	// // 单文件
	// r.POST("/upload1", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	// c.SaveUploadedFile(file,dst)
	// 	c.String(http.StatusOK, "%s uploaded!", file.Filename)
	// })
	// // 多文件
	// r.POST("upload2", func(c *gin.Context) {
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["upload[]"]

	// 	for _, file := range files {
	// 		log.Println(file.Filename)
	// 		// c.SaveUploadedFile(file,dst)
	// 	}
	// 	c.String(http.StatusOK, "%d first uploaded!", len(files))
	// })

	// //中间件
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// // 作用于单个路由
	// r.GET("/benchmark",Mylogger(),benchEndpoint)
	// // 作用于某个组
	// authorized := r.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// }

	r.Run()
}
