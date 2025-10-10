package home

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (DefaultController) Init(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	fmt.Println(username)
	ctx.String(200, "首页")
}

func (DefaultController) NewsInit(ctx *gin.Context) {
	ctx.String(200, "新闻")
}
