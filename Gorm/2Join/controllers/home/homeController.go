package home

import (
	"Gorm/2Join/models"
	"fmt"
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

func (HomeController) Test(c *gin.Context) {
	//查询数据库
	// userlist := []models.User{}

	// models.DB.Find(&userlist)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": userlist,
	// })

	//查询age大于20的用户
	userlist := []models.User{}
	models.DB.Where("age>20").Find(&userlist)

	c.JSON(200, gin.H{
		"result": userlist,
	})
}

func (HomeController) Add(c *gin.Context) {
	user := models.User{
		Id:       6,
		Username: "test gorm",
		Age:      22,
		Email:    "222@qq.com",
		AddTime:  100001111,
	}

	models.DB.Create(&user)

	fmt.Println(user)

	c.String(200, "成功增加")
}

func (HomeController) Edit(c *gin.Context) {
	//保存所有字段
	// user := models.User{Id: 5}
	// models.DB.Find(&user)
	// user.Username = "xiugaiwan"
	// user.Email = "zhanghongbin@qq.com"
	// models.DB.Save(&user)
	// fmt.Println(user)

	user := models.User{}
	models.DB.Model(&user).Where("id = ?", 5).Update("username", "xiugaier")

	c.String(200, "修改成功")
}

func (HomeController) Delete(c *gin.Context) {
	user := models.User{}
	models.DB.Where("id = ?", 6).Delete(&user)

	c.String(200, "删除成功")
}
