package main

import (
	"demo/routes"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Age int    `json:"age" xml:"age" form:"age"`
	Sex string `json:"sex" xml:"sex" form:"sex"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routes.UserRoutesInit(router)

	// router.POST("/test", func(c *gin.Context) {
	// 	name := c.DefaultQuery("name", "none")
	// 	var test Test
	// 	c.ShouldBind(&test)
	// 	fmt.Println(c.Request.Body)
	// 	// c.Status(200)                                                                        //状态行
	// 	c.Header("key", "value") //响应头
	// 	// c.Writer.WriteString(fmt.Sprintf("name:%s,age:%d,sex:%s", name, test.Age, test.Sex)) //响应体
	// })

	router.Run(":5000")
}
