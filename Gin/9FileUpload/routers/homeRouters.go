package routers

import (
	"Gin/9FileUpload/controllers/home"
	"Gin/9FileUpload/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/", middlewares.Init)
	{
		defaultRouters.GET("/", home.DefaultController{}.Init)
		defaultRouters.GET("/news", home.DefaultController{}.NewsInit)
		defaultRouters.GET("/adduser", home.DefaultController{}.AddUser)
		defaultRouters.POST("/doupload", home.DefaultController{}.DoUpload)
	}

}
