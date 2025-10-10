package article

import (
	"Gorm/1Select/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ArticleController) Index(c *gin.Context) {
	articlelist := []models.Article{}

	models.DB.Find(&articlelist)

	c.JSON(200, gin.H{
		"result": articlelist,
	})
}
