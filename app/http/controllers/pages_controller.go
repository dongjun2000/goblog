package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

// 处理静态页面
type PagesController struct {
}

// 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, 欢迎来到 Goblog！</h1>")
}

// 关于我们
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "pages.about")
}

// 404
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :( </h1><p>如有疑惑，请联系我们。</p></h1>")
}