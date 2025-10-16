package routes

import (
	"demo/controller"
	"demo/service"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.Engine) {
	service := service.NewUserService()
	usercontroller := controller.NewUserController(service)
	user := router.Group("v1")
	{
		user.POST("login", usercontroller.Login)
	}
}
