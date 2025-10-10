package main

import (
	"Gin/10Cookie/controllers/home"
	"Gin/10Cookie/models"
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
)

func AAA() {
	fmt.Println("hello world")
}

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")
	router.GET("/", home.HomeController{}.Init)
	router.GET("/news", home.HomeController{}.News)
	router.GET("/shop", home.HomeController{}.Shop)
	router.GET("/deletecookie", home.HomeController{}.DeleteCookie)

	router.Run()
}
