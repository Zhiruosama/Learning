package nav

import (
	"Gorm/1Select/models"

	"github.com/gin-gonic/gin"
)

type NavController struct{}

func (NavController) Index(c *gin.Context) {
	// navlist := []models.Nav{}

	// models.DB.Find(&navlist)

	// c.JSON(200, gin.H{
	// 	"result": navlist,
	// })

	//对查询的数据进行限定时
	//跟正常使用mysql一样使用where与and or等关键词
	//包括大于小于等于等符号同样使用
	//like、between and

	//使用Select指定返回的字段
	// navlist := []models.Nav{}
	// models.DB.Select("id,title").Find(&navlist)
	// c.JSON(200, gin.H{
	// 	"result": navlist,
	// })

	//Order排序、Limit、Offest
	// navlist := []models.Nav{}
	// models.DB.Order("id desc").Find(&navlist)
	// c.JSON(200, gin.H{
	// 	"result": navlist,
	// })
	//Limit 限制查询的数据数
	//Offest 跳过几条数据

	//用原生sql语句对数据库进行操作
	// models.DB.Exec("delete from user where id=?", 5)

	//用原生sql语句查询
	userlist := []models.User{}
	models.DB.Raw("select * from user where id > 2").Scan(&userlist)
	c.JSON(200, gin.H{
		"result": userlist,
	})
}
