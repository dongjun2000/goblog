package policies

import (
	"goblog/app/models/article"
	"goblog/pkg/auth"
)

// 是否允许修改话题
func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}