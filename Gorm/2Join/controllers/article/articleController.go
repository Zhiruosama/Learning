package article

import (
	"Gorm/2Join/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ArticleController) Index(c *gin.Context) {
	//查询文章 获取文章对应的分类
	articlelist := []models.Article{}
	models.DB.Preload("ArticleCate").Find(&articlelist)
	c.JSON(200, gin.H{
		"result": articlelist,
	})
}
