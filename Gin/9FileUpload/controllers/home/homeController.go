package home

import (
	"fmt"
	"net/http"
	"path"

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

func (DefaultController) AddUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/useradd.html", gin.H{})
}

func (DefaultController) DoUpload(ctx *gin.Context) {
	username := ctx.PostForm("username")

	file, err := ctx.FormFile("face")
	// file.Filename
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}
