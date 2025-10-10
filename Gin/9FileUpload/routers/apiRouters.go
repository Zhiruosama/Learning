package routers

import (
	"Gin/9FileUpload/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(router *gin.Engine) {

	apiRouters := router.Group("/api")
	{
		apiRouters.GET("/", api.UserListController{}.ApiInit)
		apiRouters.GET("/userlist", api.UserListController{}.UserList)
		apiRouters.GET("/article", api.ArticleControoller{}.Index)
	}
}
