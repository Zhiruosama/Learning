package controller

import (
	domain "demo/domain/dto"
	"demo/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userservice *service.UserService) *UserController {
	return &UserController{UserService: userservice}
}

func (uc *UserController) Login(c *gin.Context) {
	fmt.Println("参数校验")
	var user domain.UserDto
	c.ShouldBind(&user)

	finuser, err := uc.UserService.Login(&user)
	if err == nil {
		c.JSON(http.StatusOK, finuser)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
}
