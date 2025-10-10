package routers

import "github.com/gin-gonic/gin"

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/")
	{
		defaultRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "首页")
		})

		defaultRouters.GET("/news", func(ctx *gin.Context) {
			ctx.String(200, "news")
		})
	}

}
