package models

type ArticleCate struct {
	Id    int
	Title string
	State int
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
