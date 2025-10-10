package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (HomeController) Init(c *gin.Context) {
	//设置cookie
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, false)

	c.SetCookie("test", "短cookie", 15, "/", "localhost", false, false)

	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}

func (HomeController) News(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "cookie="+username)
}

func (HomeController) Shop(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "cookie="+username)
}

func (HomeController) DeleteCookie(c *gin.Context) {
	//删除cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(200, "删除成功")
}
