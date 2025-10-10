package routers

import (
	"Gin/6Controller/controllers/home"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/")
	{
		defaultRouters.GET("/", home.DefaultController{}.Init)
		defaultRouters.GET("/news", home.DefaultController{}.NewsInit)
	}

}
