package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Context) {
	//判断用户登录

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	//定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path: " + cCp.Request.URL.Path)
	}()
}
