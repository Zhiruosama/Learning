package main

import (
	"Gin/8Model/controllers/home"
	"Gin/8Model/models"
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

	router.Run()
}
