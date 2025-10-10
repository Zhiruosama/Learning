package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(router *gin.Engine) {

	apiRouters := router.Group("/api")
	{
		apiRouters.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "这是api页")
		})

		apiRouters.GET("/userlist", func(ctx *gin.Context) {
			ctx.String(200, "api页下的userlist")
		})
	}
}
