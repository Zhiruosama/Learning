package home

import "github.com/gin-gonic/gin"

type DefaultController struct {
}

func (DefaultController) Init(ctx *gin.Context) {
	ctx.String(200, "首页")
}

func (DefaultController) NewsInit(ctx *gin.Context) {
	ctx.String(200, "新闻")
}
