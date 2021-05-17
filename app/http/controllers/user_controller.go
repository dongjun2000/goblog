package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// 用户控制器
type UserController struct {
	BaseController
}

// 用户个人页面
func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 读取URL参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_users, err := user.Get(id)

	// 3. 如果出现错误
	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		// 4. 读取成功，显示用户文章列表
		articles, err := article.GetByUserID(_users.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}



