package routers

import (
	"Gin/7Middleware/controllers/home"
	"Gin/7Middleware/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/", middlewares.Init)
	{
		defaultRouters.GET("/", home.DefaultController{}.Init)
		defaultRouters.GET("/news", home.DefaultController{}.NewsInit)
	}

}
