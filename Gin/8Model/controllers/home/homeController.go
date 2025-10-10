package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (HomeController) Init(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}
