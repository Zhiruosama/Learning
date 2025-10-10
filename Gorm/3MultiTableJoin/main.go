package main

import (
	"Gorm/3MultiTableJoin/controllers/article"
	"Gorm/3MultiTableJoin/controllers/home"
	"Gorm/3MultiTableJoin/controllers/nav"
	"Gorm/3MultiTableJoin/models"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	router.LoadHTMLGlob("templates/**/*")

	router.Static("/static", "./static")

	router.GET("/", home.HomeController{}.Init)
	router.GET("/test", home.HomeController{}.Test)
	router.GET("/add", home.HomeController{}.Add)
	router.GET("/edit", home.HomeController{}.Edit)
	router.GET("/delete", home.HomeController{}.Delete)

	router.GET("/nav/index", nav.NavController{}.Index)

	router.GET("/article/index", article.ArticleController{}.Index)

	router.Run()
}
