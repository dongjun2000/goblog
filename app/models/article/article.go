package article

import (
	"goblog/app/models"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body string `gorm:"type:text;not null;" valid:"body"`
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}