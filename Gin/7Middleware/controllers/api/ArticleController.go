package api

import "github.com/gin-gonic/gin"

type ArticleControoller struct {
}

func (ArticleControoller) Index(c *gin.Context) {
	c.String(200, "文章列表")
}
