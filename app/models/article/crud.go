package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}